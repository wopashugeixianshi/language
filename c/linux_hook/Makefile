test: main.c hook.so
	gcc -o test main.c

hook.so: hook.c
	gcc -fPIC -shared -o libhook.so hook.c -ldl
	#gcc -fpic -shared -o libhook.so hook.c -ldl
