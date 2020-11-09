# donut
Port of the spinning donut from C to Go\
Donut is simply a framebuffer and a Z-buffer which the pixels are rendered onto. Since it is rendering relatively low-resolution ASCII art, all that needs to be done is to plot pixels along the surface of the torus at fixed-angle increments. This is done densely enough that the final result looks solid. The “pixels” it plots are ASCII characters corresponding to the illumination value of the surface at each point: .,-~:;=!*#$@ from dimmest to brightest. No raytracing required.\
Originial in C: https://www.a1k0n.net/2006/09/15/obfuscated-c-donut.html
