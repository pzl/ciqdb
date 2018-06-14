#!/usr/bin/env python

import sys, os
import xml.etree.ElementTree as ET
from enum import Enum
from collections import namedtuple


def to_int(bytestring):
	return int.from_bytes(bytestring,'big')


class EntryPoint(namedtuple('EntryPoint', 'uuid, module, symbol, label, icons, apptype')):
	__slots__ = ()
	def __str__(self):
		return "UUID: {}\nType: {}\n{}: {}".format(self.uuid,self.apptype,self.label,self.symbol)
		
class AppType(Enum):
	WATCH_FACE = 0
	APP = 1
	DATA_FIELD = 2
	WIDGET = 3
	BACKGROUND = 4
	AUDIO = 5

	def __str__(self):
		names = ["watch face", "app", "data field", "widget", "background", "audio provider"]
		if self.value < len(names):
			return names[self.value]
		else:
			return "unknown"

class App(object):
	"""Represents a Connect IQ app from a .prg"""

	def __init__(self, filename):
		super(App, self).__init__()
		self.filename = filename
		self.entry_points = PrgParser(self.filename).parse()
		
	def __str__(self):
		for e in self.entry_points:
			return str(e)
		

class PrgParser(object):
	"""Reads a .prg file"""

	entry_point_header=b'\x60\x60\xc0\xde'

	def __init__(self, filename):
		super(PrgParser, self).__init__()
		self.filename = filename
		
	def parse(self):
		entry_points=[]
		with open(self.filename,"rb") as f:
			while True:
				sec = self.parse_section(f)
				if sec is None: continue
				if sec is False: break
				entry_points = sec
				break
		return entry_points

	def parse_section(self,f):
		header = f.read(8)
		if not header: return False

		sec_type,length = header[:4],to_int(header[4:])
		if sec_type == self.entry_point_header:
			return self.parse_entries(f)
		else:
			f.seek(length,1) # skip section, move through it

	def parse_entries(self,f):
		entries = []
		num_entries = to_int(f.read(2))

		for _ in range(num_entries):
			entries.append(self.parse_entry(f))

		return entries


	def parse_entry(self,f):
		data = f.read(36)
		uuid = data[:16].hex()
		module = to_int(data[16:20])
		symbol = to_int(data[20:24]) # see .prg.debug : <symbolTable><entry id="X"/>
		label = to_int(data[24:28])
		icons = to_int(data[28:32])
		apptype = AppType(to_int(data[32:36]))

		label = self.get_symbol(label)
		symbol = self.get_symbol(symbol)

		return EntryPoint(uuid,module,symbol,label,icons,apptype)

	def get_symbol(self,s):
		try:
			et = ET.parse(os.path.expanduser(self.filename+".debug.xml"))
		except FileNotFoundError:
			return s
		root = et.getroot()
		symbol = root.find('./symbolTable/entry[@id="{}"]'.format(s,))
		return symbol.get("symbol") if symbol is not None else s

def main():
	import sys
	if len(sys.argv) < 2:
		print("argument required: path to .prg")
		sys.exit(1)
	a = App(sys.argv[1])
	print(a)

if __name__ == '__main__':
	main()