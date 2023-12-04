package handler

import (
	"strconv"

	"github.com/IvanTarjan/final-go-g5/internal/domain"
	"github.com/IvanTarjan/final-go-g5/internal/patient"
	"github.com/IvanTarjan/final-go-g5/pkg/web"
	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	service patient.ServicePatients
}

func NewPatientHandler(service *patient.ServicePatients) *PatientHandler {
	return &PatientHandler{service: *service}
}

// HandlerCreate is a function that creates a patient
// @Summary patient example
// @Description Create a new patient
// @Tags patient
// @Param patient body domain.Patient true "New Patient"
// @Param token header string true "Token for authentication"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /patient [post]
func (h *PatientHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var patientTemp domain.Patient
		if err := ctx.BindJSON(&patientTemp); err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		createdPatient, err := h.service.Create(patientTemp)
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 201, createdPatient)
	}
}

// HandlerGetAll retrieves all patients
// @Summary Retrieve all patients
// @Description Get a list of all patients
// @Tags patient
// @Produce json
// @Success 200 {array} domain.Patient "List of Patients"
// @Failure 400 {object} web.errorResponse "Bad Request"
// @Router /patients [get]
func (h *PatientHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		patientList, err := h.service.GetAll()
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, patientList)
	}
}

// HandlerGetById retrieves a patient by id
// @Summary Retrieve a patient by id
// @Description Get a single patient by their id
// @Tags patient
// @Produce json
// @Param id path int true "Patient ID"
// @Success 200 {object} domain.Patient "Patient Data"
// @Failure 400 {object} web.errorResponse "Bad Request"
// @Router /patient/{id} [get]
func (h *PatientHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		patientTemp, err := h.service.GetByID(int64(id))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, patientTemp)
	}
}

// HandlerGetById retrieves a patient by id
// @Summary Retrieve a patient by id
// @Description Get a single patient by their id
// @Tags patient
// @Produce json
// @Param id path int true "Patient ID"
// @Param token header string true "Token for authentication"
// @Success 200 {object} domain.Patient "Patient Data"
// @Failure 400 {object} web.errorResponse "Bad Request"
// @Router /patient/{id} [get]
func (h *PatientHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		var patientTemp domain.Patient
		if err := ctx.BindJSON(&patientTemp); err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		updatedPatient, err := h.service.Update(patientTemp, int64(id))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, updatedPatient)
	}
}

// HandlerPatch partially updates a patient's information
// @Summary Partially update a patient
// @Description Partially update a patient's information by their id
// @Tags patient
// @Accept json
// @Produce json
// @Param id path int true "Patient ID"
// @Param token header string true "Token for authentication"
// @Param patient body domain.Patient true "Partial Patient Data"
// @Success 200 {object} domain.Patient "Patched Patient Data"
// @Failure 400 {object} web.errorResponse "Bad Request"
// @Router /patient/{id} [patch]
func (h *PatientHandler) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		var patientTemp domain.Patient
		if err := ctx.BindJSON(&patientTemp); err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		patchedPatient, err := h.service.Patch(patientTemp, int64(id))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}

		web.Success(ctx, 200, patchedPatient)
	}
}

// HandlerDelete deletes a patient
// @Summary Delete a patient
// @Description Delete a patient by their id
// @Tags patient
// @Produce json
// @Param id path int true "Patient ID"
// @Param token header string true "Token for authentication"
// @Success 200 {string} string "Deleted patient with id {id}"
// @Failure 400 {object} web.errorResponse "Bad Request"
// @Router /patient/{id} [delete]
func (h *PatientHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		err = h.service.Delete(int64(id))
		if err != nil {
			web.Failure(ctx, 400, err.Error())
			return
		}
		web.Success(ctx, 200, "Deleted patient with id "+strconv.FormatInt(int64(id), 10))
	}
}
