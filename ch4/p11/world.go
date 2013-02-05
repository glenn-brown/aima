package p11

import (
	"github.com/banthar/gl"
	"github.com/glenn-brown/vu"
	"math/rand"
)

type World struct {
	rng       *rand.Rand
	goal      vu.Point
	robot     vu.Point
	obstacles vu.Polygons
	budget    float64
	percept   vu.Points
}

// Create a world.  
func NewWorld(obstacles int) *World {
	w := &World{
		rng:    rand.New(rand.NewSource(123456789)),
		budget: 1000.0,
	}
	for i := 0; i < obstacles; i++ {
		w.obstacles = append(w.obstacles, w.newObstacle())
	}
	w.goal = w.randomPoint()
	w.robot = w.randomPoint()
	return w
}

// Return a random point that is not inside an obstacle.
func (w *World) randomPoint() vu.Point {
outer:
	for {
		p := vu.Point{w.rng.Float64(), w.rng.Float64()}
		for _, poly := range w.obstacles {
			if poly.Contains(p) {
				continue outer
			}
		}
		return p
	}
	panic("Unreachable.")
}

// Create a new obstacle that does not collide with any other.
func (w *World) newObstacle() vu.Polygon {
	obstacle := w.grow(w.grow(vu.Polygon{w.randomPoint()}))
	// obstacle = w.grow(obstacle)
	// obstacle = w.grow(obstacle)
	// obstacle = w.grow(obstacle)
	return obstacle
}

// Grow an obstacle such that it remains convex and does not overlay any other.
func (w *World) grow(poly vu.Polygon) vu.Polygon {
	last := len(poly)
	nu := append(poly, w.randomPoint())
	for {
		if nu.IsConvex() &&
			!w.obstacles.Intersect(vu.Segment{nu[last-1], nu[last]}) &&
			!w.obstacles.Intersect(vu.Segment{nu[last], nu[0]}) {
			return nu
		}
		nu[last] = w.randomPoint()
	}
	panic("unreachable")
}

func (w *World) Step(agent Agent) {

	// Let agent request the next move.

	move := agent.Move(w.goal.Sub(w.robot), w.percept)
	request := vu.Segment{w.robot, w.robot.Add(move)}

	// Compute stop location.

	end := request.B // Speculative
	collision := w.obstacles.Intersection(request)
	if collision != nil {
		end = *collision
	}

	// Perform the move, charging for it.

	w.robot = end
	w.budget -= vu.Segment{w.robot, end}.Len()

	// Teleport if budget was exceeded or stopped at the goal.

	if w.budget <= 0 || (move.Equals(vu.Point{0, 0}) && w.robot.Equals(w.goal)) {
		w.robot = w.randomPoint()
		w.budget = 1000.0
	}

	// Compute the relative positions of vertices visible to the robot.

	w.percept = vu.Points{}
	for _, pp := range w.obstacles {
		for _, pt := range pp {
			x := w.obstacles.Intersection(vu.Segment{w.robot, pt})
			if x == nil {
				w.percept = append(w.percept, pt.Sub(w.robot))
			}
		}
	}
}

func (world *World) Render(w, h, d float64) {
	gl.PushMatrix()
	gl.Scaled(w, h, d)

	// Draw the goal
	gl.Begin(gl.POINTS)
	gl.Color3ub(255, 0, 0)
	gl.PointSize(6)
	gl.Vertex2d(world.goal.X, world.goal.Y)
	gl.End()

	// Draw the obstacles
	gl.Color4ub(0, 0, 255, 63)
	world.obstacles.Render()

	// Draw the percept
	gl.Color4ub(0, 255, 0, 127)
	gl.LineWidth(1.5)
	gl.Begin(gl.LINES)
	for _, pt := range world.percept {
		gl.Vertex2d(world.robot.X, world.robot.Y)
		gl.Vertex2d(world.robot.X+pt.X, world.robot.Y+pt.Y)
	}
	gl.End()

	// Draw the map
	gl.Color3ub(0, 0, 255)
	gl.LineWidth(2)
	// FIXME

	// Draw the robot
	gl.Color3ub(0, 255, 0)
	gl.PointSize(4)
	gl.Begin(gl.POINTS)
	gl.Vertex2d(world.robot.X, world.robot.Y)
	gl.End()

	gl.PopMatrix()
}

type Agent interface {
	Render(w, h, d float64)
	Move(goal vu.Point, percept vu.Points) vu.Point
}
