package controllers

import (
	"Apps-I_Desa_Backend/dtos"
	"Apps-I_Desa_Backend/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type SubDimensionController struct {
	subDimensionService *services.SubDimensionService
	validate            *validator.Validate
}

func NewSubDimensionController(
	subDimensionService *services.SubDimensionService,
) *SubDimensionController {
	return &SubDimensionController{
		subDimensionService: subDimensionService,
		validate:            validator.New(),
	}
}

func (c *SubDimensionController) CreateSubDimensionPendidikan(ctx *fiber.Ctx) error {
	var request dtos.AddSubDimensionPendidikanRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	response, err := c.subDimensionService.CreateSubDimensionPendidikan(&request, ctx)
	if err != nil {
		if err.Error() == "village ID not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Village ID not found",
				"error":   "Check your token",
			})
		} else if err.Error() == "invalid village ID" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid village ID",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to create sub dimension pendidikan" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Invalid sub-dimension type",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to commit transaction" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to commit transaction",
				"error":   err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create sub-dimension pendidikan",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *SubDimensionController) CreateSubDimensionKesehatan(ctx *fiber.Ctx) error {
	var request dtos.AddSubDimensionKesehatanRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	response, err := c.subDimensionService.CreateSubDimensionKesehatan(&request, ctx)
	if err != nil {
		if err.Error() == "village ID not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Village ID not found",
				"error":   "Check your token",
			})
		} else if err.Error() == "invalid village ID" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid village ID",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to create sub dimension kesehatan" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Invalid sub-dimension type",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to commit transaction" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to commit transaction",
				"error":   err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create sub-dimension kesehatan",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *SubDimensionController) CreateSubDimensionUtilitasDasar(ctx *fiber.Ctx) error {
	var request dtos.AddSubDimensionUtilitasDasarRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	response, err := c.subDimensionService.CreateSubDimensionUtilitasDasar(&request, ctx)
	if err != nil {
		if err.Error() == "village ID not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Village ID not found",
				"error":   "Check your token",
			})
		} else if err.Error() == "invalid village ID" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid village ID",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to create sub dimension utilitas dasar" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Invalid sub-dimension type",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to commit transaction" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to commit transaction",
				"error":   err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create sub-dimension utilitas dasar",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *SubDimensionController) CreateSubDimensionAktivitas(ctx *fiber.Ctx) error {
	var request dtos.AddSubDimensionAktivitasRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	response, err := c.subDimensionService.CreateSubDimensionAktivitas(&request, ctx)
	if err != nil {
		if err.Error() == "village ID not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Village ID not found",
				"error":   "Check your token",
			})
		} else if err.Error() == "invalid village ID" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid village ID",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to create sub dimension aktivitas" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Invalid sub-dimension type",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to commit transaction" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to commit transaction",
				"error":   err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create sub-dimension aktivitas",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *SubDimensionController) CreateSubDimensionFasilitasMasyarakat(ctx *fiber.Ctx) error {
	var request dtos.AddSubDimensionFasilitasMasyarakatRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	response, err := c.subDimensionService.CreateSubDimensionFasilitasMasyarakat(&request, ctx)
	if err != nil {
		if err.Error() == "village ID not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Village ID not found",
				"error":   "Check your token",
			})
		} else if err.Error() == "invalid village ID" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid village ID",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to create sub dimension fasilitas masyarakat" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Invalid sub-dimension type",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to commit transaction" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to commit transaction",
				"error":   err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create sub-dimension fasilitas masyarakat",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *SubDimensionController) CreateSubDimensionProduksiDesa(ctx *fiber.Ctx) error {
	var request dtos.AddSubDimensionProduksiDesaRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	response, err := c.subDimensionService.CreateSubDimensionProduksiDesa(&request, ctx)
	if err != nil {
		if err.Error() == "village ID not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Village ID not found",
				"error":   "Check your token",
			})
		} else if err.Error() == "invalid village ID" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid village ID",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to create sub dimension produksi desa" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Invalid sub-dimension type",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to commit transaction" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to commit transaction",
				"error":   err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create sub-dimension produksi desa",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *SubDimensionController) CreateSubDimensionFasilitasPendukungEkonomi(ctx *fiber.Ctx) error {
	var request dtos.AddSubDimensionFasilitasPendukungEkonomiRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	response, err := c.subDimensionService.CreateSubDimensionFasilitasPendukungEkonomi(
		&request,
		ctx,
	)
	if err != nil {
		if err.Error() == "village ID not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Village ID not found",
				"error":   "Check your token",
			})
		} else if err.Error() == "invalid village ID" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid village ID",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to create sub dimension fasilitas pendukung ekonomi" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Invalid sub-dimension type",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to commit transaction" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to commit transaction",
				"error":   err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create sub-dimension fasilitas pendukung ekonomi",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *SubDimensionController) CreateSubDimensionPengelolaanLingkungan(ctx *fiber.Ctx) error {
	var request dtos.AddSubDimensionPengelolaanLingkunganRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	response, err := c.subDimensionService.CreateSubDimensionPengelolaanLingkungan(&request, ctx)
	if err != nil {
		if err.Error() == "village ID not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Village ID not found",
				"error":   "Check your token",
			})
		} else if err.Error() == "invalid village ID" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid village ID",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to create sub dimension pengelolaan lingkungan" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Invalid sub-dimension type",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to commit transaction" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to commit transaction",
				"error":   err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create sub-dimension pengelolaan lingkungan",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *SubDimensionController) CreateSubDimensionPenanggulanganBencana(ctx *fiber.Ctx) error {
	var request dtos.AddSubDimensionPenanggulanganBencanaRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	response, err := c.subDimensionService.CreateSubDimensionPenanggulanganBencana(&request, ctx)
	if err != nil {
		if err.Error() == "village ID not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Village ID not found",
				"error":   "Check your token",
			})
		} else if err.Error() == "invalid village ID" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid village ID",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to create sub dimension penanggulangan bencana" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Invalid sub-dimension type",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to commit transaction" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to commit transaction",
				"error":   err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create sub-dimension penanggulangan bencana",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *SubDimensionController) CreateSubDimensionKondisiAksesJalan(ctx *fiber.Ctx) error {
	var request dtos.AddSubDimensionKondisiAksesJalanRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	response, err := c.subDimensionService.CreateSubDimensionKondisiAksesJalan(&request, ctx)
	if err != nil {
		if err.Error() == "village ID not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Village ID not found",
				"error":   "Check your token",
			})
		} else if err.Error() == "invalid village ID" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid village ID",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to create sub dimension kondisi akses jalan" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Invalid sub-dimension type",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to commit transaction" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to commit transaction",
				"error":   err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create sub-dimension kondisi akses jalan",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *SubDimensionController) CreateSubDimensionKemudahanAkses(ctx *fiber.Ctx) error {
	var request dtos.AddSubDimensionKemudahanAksesRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	response, err := c.subDimensionService.CreateSubDimensionKemudahanAkses(&request, ctx)
	if err != nil {
		if err.Error() == "village ID not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Village ID not found",
				"error":   "Check your token",
			})
		} else if err.Error() == "invalid village ID" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid village ID",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to create sub dimension kemudahan akses" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Invalid sub-dimension type",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to commit transaction" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to commit transaction",
				"error":   err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create sub-dimension kemudahan akses",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *SubDimensionController) CreateSubDimensionKelembagaanPelayananDesa(ctx *fiber.Ctx) error {
	var request dtos.AddSubDimensionKelembagaanPelayananDesaRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	response, err := c.subDimensionService.CreateSubDimensionKelembagaanPelayananDesa(&request, ctx)
	if err != nil {
		if err.Error() == "village ID not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Village ID not found",
				"error":   "Check your token",
			})
		} else if err.Error() == "invalid village ID" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid village ID",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to create sub dimension kelembagaan pelayanan desa" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Invalid sub-dimension type",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to commit transaction" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to commit transaction",
				"error":   err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create sub-dimension kelembagaan pelayanan desa",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *SubDimensionController) CreateSubDimensionTataKelolaKeuanganDesa(ctx *fiber.Ctx) error {
	var request dtos.AddSubDimensionTataKelolaKeuanganDesaRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	response, err := c.subDimensionService.CreateSubDimensionTataKelolaKeuanganDesa(&request, ctx)
	if err != nil {
		if err.Error() == "village ID not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Village ID not found",
				"error":   "Check your token",
			})
		} else if err.Error() == "invalid village ID" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid village ID",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to create sub dimension tata kelola keuangan desa" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Invalid sub-dimension type",
				"error":   err.Error(),
			})
		} else if err.Error() == "failed to commit transaction" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to commit transaction",
				"error":   err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create sub-dimension tata kelola keuangan desa",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}
