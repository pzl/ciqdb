import prg
import shell as Shell
import sys
import os


def main():
	if len(sys.argv) < 2:
		print("argument required: .prg file to debug")
		sys.exit(1)

	filename = os.path.basename(sys.argv[1])

	app = prg.App(sys.argv[1])
	print(app)

	uuid = app.entry_points[0].uuid

	shell = Shell.Shell("/home/dan/dev/connectiq/connectiq-sdk-lin-3.0.0-beta1")
	shell.connect()

	prg_path = "0:/GARMIN/APPS/"
	if app.entry_points[0].apptype == prg.AppType.AUDIO:
		prg_path += "MEDIA/"
	prg_path += filename

	shell.write("push {} {}".format(sys.argv[1],prg_path)) 
	info = shell.read(5)
	while info is not None:
		if info:
			sys.stdout.write(info)
			sys.stdout.flush()
		info = shell.read(5)

	shell.write("ciq")
	print(shell.readall(),flush=True)

	shell.write("[1][0]openDevice {}".format("fr935")) # @TODO: device name
	print(shell.readall(),flush=True)

	shell.write("[2][0]debugApp {}".format(uuid))
	print(shell.readall(),flush=True)

	shell.write("[3][{}][1]RE".format(uuid))

	
	#@todo: read

	#@todo: breakpoints and all that



if __name__ == '__main__':
	main()