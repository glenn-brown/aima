/* Public-domain gears demo from the SDL web site, translated to go. */
package main

import (
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
	"github.com/banthar/gl"
	"math"
)

var t0 GLint = 0
var frames GLint = 0

/**

  Draw a gear wheel.  You'll probably want to call this function when
  building a display list since we do a lot of trig here.

  Input:  inner_radius - radius of hole at center
          outer_radius - radius at center of teeth
          width - width of gear
          teeth - number of teeth
          tooth_depth - depth of tooth

 **/

func gear(
	inner_radius, outer_radius, width GLfloat,
	teeth GLint, tooth_depth GLfloat) {

	r0 := inner_radius
	r1 := outer_radius - tooth_depth/2.0
	r2 := outer_radius + tooth_depth/2.0

	da := 2.0 * math.Pi / GLfloat(teeth) / 4.0

	glShadeModel(GL_FLAT)

	glNormal3f(0.0, 0.0, 1.0)

	/* draw front face */
	glBegin(GL_QUAD_STRIP)
	for i := 0; i <= teeth; i++ {
		angle := i * 2.0 * M_PI / teeth
		glVertex3f(r0*cos(angle), r0*sin(angle), width*0.5)
		glVertex3f(r1*cos(angle), r1*sin(angle), width*0.5)
		if i < teeth {
			glVertex3f(r0*cos(angle), r0*sin(angle), width*0.5)
			glVertex3f(r1*cos(angle+3*da), r1*sin(angle+3*da), width*0.5)
		}
	}
	glEnd()

	/* draw front sides of teeth */
	glBegin(GL_QUADS)
	da = 2.0 * M_PI / teeth / 4.0
	for i := 0; i < teeth; i++ {
		angle = i * 2.0 * M_PI / teeth

		glVertex3f(r1*cos(angle), r1*sin(angle), width*0.5)
		glVertex3f(r2*cos(angle+da), r2*sin(angle+da), width*0.5)
		glVertex3f(r2*cos(angle+2*da), r2*sin(angle+2*da), width*0.5)
		glVertex3f(r1*cos(angle+3*da), r1*sin(angle+3*da), width*0.5)
	}
	glEnd()

	glNormal3f(0.0, 0.0, -1.0)

	/* draw back face */
	glBegin(GL_QUAD_STRIP)
	for i := 0; i <= teeth; i++ {
		angle := i * 2.0 * M_PI / teeth
		glVertex3f(r1*cos(angle), r1*sin(angle), -width*0.5)
		glVertex3f(r0*cos(angle), r0*sin(angle), -width*0.5)
		if i < teeth {
			glVertex3f(r1*cos(angle+3*da), r1*sin(angle+3*da), -width*0.5)
			glVertex3f(r0*cos(angle), r0*sin(angle), -width*0.5)
		}
	}
	glEnd()

	/* draw back sides of teeth */
	glBegin(GL_QUADS)
	da = 2.0 * M_PI / teeth / 4.0
	for i = 0; i < teeth; i++ {
		angle := i * 2.0 * M_PI / teeth

		glVertex3f(r1*cos(angle+3*da), r1*sin(angle+3*da), -width*0.5)
		glVertex3f(r2*cos(angle+2*da), r2*sin(angle+2*da), -width*0.5)
		glVertex3f(r2*cos(angle+da), r2*sin(angle+da), -width*0.5)
		glVertex3f(r1*cos(angle), r1*sin(angle), -width*0.5)
	}
	glEnd()

	/* draw outward faces of teeth */
	glBegin(GL_QUAD_STRIP)
	for i := 0; i < teeth; i++ {
		angle = i * 2.0 * M_PI / teeth

		glVertex3f(r1*cos(angle), r1*sin(angle), width*0.5)
		glVertex3f(r1*cos(angle), r1*sin(angle), -width*0.5)
		u = r2*cos(angle+da) - r1*cos(angle)
		v = r2*sin(angle+da) - r1*sin(angle)
		len = sqrt(u*u + v*v)
		u /= len
		v /= len
		glNormal3f(v, -u, 0.0)
		glVertex3f(r2*cos(angle+da), r2*sin(angle+da), width*0.5)
		glVertex3f(r2*cos(angle+da), r2*sin(angle+da), -width*0.5)
		glNormal3f(cos(angle), sin(angle), 0.0)
		glVertex3f(r2*cos(angle+2*da), r2*sin(angle+2*da), width*0.5)
		glVertex3f(r2*cos(angle+2*da), r2*sin(angle+2*da), -width*0.5)
		u = r1*cos(angle+3*da) - r2*cos(angle+2*da)
		v = r1*sin(angle+3*da) - r2*sin(angle+2*da)
		glNormal3f(v, -u, 0.0)
		glVertex3f(r1*cos(angle+3*da), r1*sin(angle+3*da), width*0.5)
		glVertex3f(r1*cos(angle+3*da), r1*sin(angle+3*da), -width*0.5)
		glNormal3f(cos(angle), sin(angle), 0.0)
	}

	glVertex3f(r1*cos(0), r1*sin(0), width*0.5)
	glVertex3f(r1*cos(0), r1*sin(0), -width*0.5)

	glEnd()

	glShadeModel(GL_SMOOTH)

	/* draw inside radius cylinder */
	glBegin(GL_QUAD_STRIP)
	for i := 0; i <= teeth; i++ {
		angle = i * 2.0 * M_PI / teeth
		glNormal3f(-cos(angle), -sin(angle), 0.0)
		glVertex3f(r0*cos(angle), r0*sin(angle), -width*0.5)
		glVertex3f(r0*cos(angle), r0*sin(angle), width*0.5)
	}
	glEnd()

}

var view_rotx, view_roty, view_rotz GLfloat = 20.0, 30.0, 0.0
var GLfloat angle = 0.0

func draw() {
	glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)

	glPushMatrix()
	glRotatef(view_rotx, 1.0, 0.0, 0.0)
	glRotatef(view_roty, 0.0, 1.0, 0.0)
	glRotatef(view_rotz, 0.0, 0.0, 1.0)

	glPushMatrix()
	glTranslatef(-3.0, -2.0, 0.0)
	glRotatef(angle, 0.0, 0.0, 1.0)
	glCallList(gear1)
	glPopMatrix()

	glPushMatrix()
	glTranslatef(3.1, -2.0, 0.0)
	glRotatef(-2.0*angle-9.0, 0.0, 0.0, 1.0)
	glCallList(gear2)
	glPopMatrix()

	glPushMatrix()
	glTranslatef(-3.1, 4.2, 0.0)
	glRotatef(-2.0*angle-25.0, 0.0, 0.0, 1.0)
	glCallList(gear3)
	glPopMatrix()

	glPopMatrix()

	sdl.GL_SwapBuffers()

	frames++
	{
		t := sdl.GetTicks()
		if t-t0 >= 5000 {
			seconds := (t - t0) / 1000.0
			fps := frames / seconds
			printf("%d frames in %g seconds = %g FPS\n", frames, seconds, fps)
			t0 = t
			frames = 0
		}
	}
}

func idle() {
	angle += 2.0
}

/* new window size or exposure */
func reshape(width, height int) {
	h := GLfloat(height) / GLfloat(width)

	glViewport(0, 0, GLint(width), GLint(height))
	glMatrixMode(GL_PROJECTION)
	glLoadIdentity()
	glFrustum(-1.0, 1.0, -h, h, 5.0, 60.0)
	glMatrixMode(GL_MODELVIEW)
	glLoadIdentity()
	glTranslatef(0.0, 0.0, -40.0)
}

func init(argc int, argv []string) {
	pos := [4]GLfloat{5.0, 5.0, 10.0, 0.0}
	red := [4]GLfloat{0.8, 0.1, 0.0, 1.0}
	green := [4]GLfloat{0.0, 0.8, 0.2, 1.0}
	blue := [4]GLfloat{0.2, 0.2, 1.0, 1.0}

	glLightfv(GL_LIGHT0, GL_POSITION, pos)
	glEnable(GL_CULL_FACE)
	glEnable(GL_LIGHTING)
	glEnable(GL_LIGHT0)
	glEnable(GL_DEPTH_TEST)

	/* make the gears */
	gear1 = glGenLists(1)
	glNewList(gear1, GL_COMPILE)
	glMaterialfv(GL_FRONT, GL_AMBIENT_AND_DIFFUSE, red)
	gear(1.0, 4.0, 1.0, 20, 0.7)
	glEndList()

	gear2 = glGenLists(1)
	glNewList(gear2, GL_COMPILE)
	glMaterialfv(GL_FRONT, GL_AMBIENT_AND_DIFFUSE, green)
	gear(0.5, 2.0, 2.0, 10, 0.7)
	glEndList()

	gear3 = glGenLists(1)
	glNewList(gear3, GL_COMPILE)
	glMaterialfv(GL_FRONT, GL_AMBIENT_AND_DIFFUSE, blue)
	gear(1.3, 2.0, 0.5, 10, 0.7)
	glEndList()

	glEnable(GL_NORMALIZE)

	if argc > 1 && strcmp(argv[1], "-info") == 0 {
		printf("GL_RENDERER   = %s\n", glGetString(GL_RENDERER))
		printf("GL_VERSION    = %s\n", glGetString(GL_VERSION))
		printf("GL_VENDOR     = %s\n", glGetString(GL_VENDOR))
		printf("GL_EXTENSIONS = %s\n", glGetString(GL_EXTENSIONS))
	}
}

func main() {

	sdl.Init(sdl.INIT_VIDEO)

	screen = sdl.SetVideoMode(300, 300, 16, sdl.OPENGL|sdl.RESIZABLE)
	if !screen {
		fprintf(stderr, "Couldn't set 300x300 GL video mode: %s\n", sdl.GetError())
		sdl.Quit()
		exit(2)
	}
	sdl.WM_SetCaption("Gears", "gears")

	init(argc, argv)
	reshape(screen.w, screen.h)
	done := false
	for !done {
		var event sdl.Event

		idle()
		for sdl.PollEvent(&event) {
			switch event.Type {
			case sdl.VIDEORESIZE:
				screen = sdl.SetVideoMode(event.resize.w, event.resize.h, 16,
					sdl.OPENGL|sdl.RESIZABLE)
				if screen {
					reshape(screen.w, screen.h)
				} else {
					/* Uh oh, we couldn't set the new video mode?? */
				}
				break

			case sdl.QUIT:
				done = true
				break
			}
		}
		keys = sdl.GetKeyState(NULL)

		if keys[SDLK_ESCAPE] {
			done = true
		}
		if keys[SDLK_UP] {
			view_rotx += 5.0
		}
		if keys[SDLK_DOWN] {
			view_rotx -= 5.0
		}
		if keys[SDLK_LEFT] {
			view_roty += 5.0
		}
		if keys[SDLK_RIGHT] {
			view_roty -= 5.0
		}
		if keys[SDLK_z] {
			if sdl.GetModState() & KMOD_SHIFT {
				view_rotz -= 5.0
			} else {
				view_rotz += 5.0
			}
		}

		draw()
	}
	sdl.Quit()
	return 0 /* ANSI C requires main to return int. */
}
