package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct{}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request) {
	paramsWorkoutID := chi.URLParam(r, "id")
	if paramsWorkoutID == "" {
		http.NotFound(w, r)
		return
	}

	// convert string id to base 10, 64 bit int
	workoutID, err := strconv.ParseInt(paramsWorkoutID, 10, 64)
	if err != nil {
		//TODO: - change this at a later stage
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "This is the id of the workout: %d\n", workoutID)

}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Workout created")
}
