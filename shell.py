#!/usr/bin/env python

import os
import sys
import subprocess
from threading import Thread
from enum import Enum, auto

try:
	from queue import Queue, Empty
except ImportError:
	from Queue import Queue, Empty # python 2


class Command(Enum):
	BR=auto() # add breakpoint
	CL=auto() # clear breakpoints
	DE=auto() # delete breakpoint
	DO=auto() # get var details
	BT=auto() # get backtrace
	EB=auto() # get breakpoints
	GA=auto() # get item in array
	GL=auto() # get local value
	GV=auto() # get object content
	GP=auto() # get program counter
	KI=auto() # kill program
	SP=auto() # pause execution
	RE=auto() # resume execution
	ST=auto() # step instruction

class ObjectType(Enum):
	NULL = 0
	INT = 1
	FLOAT = 2
	STRING = 3
	OBJECT = 4
	ARRAY = 5
	METHOD = 6
	CLASSDEF = 7
	SYMBOL = 8
	BOOLEAN = 9
	MODULEDEF = 10
	HASH = 11
	RESOURCE = 12 # then BITMAP=0, FONT=1
	PRIMITIVE_OBJECT = 13
	LONG = 14
	DOUBLE = 15
	WEAK_POINTER = 16
	PRIMITIVE_MODULE = 17
	SYSTEM_POINTER = 18
	CHAR = 19


class Shell(object):
	"""communicator with the shell bin"""
	def __init__(self, sdk_path):
		super(Shell, self).__init__()
		self.sdk_path = sdk_path

		self._q = Queue()
	



	def connect(self):
		cmd = os.path.join(self.sdk_path,"bin","shell")
		self.proc = subprocess.Popen(cmd, shell=True,stdin=subprocess.PIPE,stdout=subprocess.PIPE,stderr=subprocess.PIPE,bufsize=1,universal_newlines=True)

		def reader(stream, queue):
			while True:
				char = stream.read(1)
				if char:
					queue.put(char)
				else:
					raise UnexpectedEndOfStream

		self._reader_thread = Thread(target=reader, args=(self.proc.stdout,self._q))
		self._reader_thread.daemon = True
		self._reader_thread.start()

		startup_prinout = self.readall()
		print(startup_prinout,end="")

	def _exitcheck(self):
		if self.proc.poll() is not None:
			print("shell exited")
			sys.exit(0)

	def write(self,data):
		self._exitcheck()
		self.proc.stdin.write(data+"\n")
		sys.stdout.write(data+"\n")
		sys.stdout.flush()

	def read(self, timeout=1):
		self._exitcheck()
		try:
			return self._q.get(block=timeout is not None, timeout=timeout)
		except Empty:
			return None

	def readall(self):
		s=""
		info = self.read()
		while info is not None:
			if info:
				s += info
			info = self.read()
		return s


def main():

	shell = Shell("/home/dan/dev/connectiq/connectiq-sdk-lin-3.0.0-beta1")
	shell.connect()

	shell.write("help push")

	
	info = shell.read()
	while info is not None:
		if info:
			sys.stdout.write(info)
			sys.stdout.flush()
		info = shell.read()

class UnexpectedEndOfStream(Exception): pass

if __name__ == '__main__':
	main()