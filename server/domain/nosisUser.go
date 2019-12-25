package domain

import (
	"time"
)

type NosisUser struct {
	Id          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Raw         string `gorm:"type:text;;not null" json:"raw"`
	CandidateId uint64
	Candidate   Candidate `gorm:"foreignkey:id" json:"ignore"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Variable struct {
	Nombre      string `json:"Nombre"`
	Valor       string `json:"Valor"`
	Descripcion string `json:"Descripcion"`
	Tipo        string `json:"Tipo"`
	FechaAct    string `json:"FechaAct,omitempty"`
	Detalle     string `json:"Detalle,omitempty"`
}

type Datos struct {
	Variables []Variable `json:"Variables"`
}

type NosisResponse struct {
	Contenido struct {
		Pedido struct {
			Usuario   int    `json:"Usuario"`
			Documento string `json:"Documento"`
			VR        string `json:"VR"`
			Timeout   int    `json:"Timeout"`
		} `json:"Pedido"`
		Resultado struct {
			Estado         int    `json:"Estado"`
			Novedad        string `json:"Novedad"`
			Tiempo         int    `json:"Tiempo"`
			FechaRecepcion string `json:"FechaRecepcion"`
			Transaccion    string `json:"Transaccion"`
			Referencia     string `json:"Referencia"`
			Servidor       string `json:"Servidor"`
			Version        string `json:"Version"`
		} `json:"Resultado"`
		Datos Datos `json:"Datos"`
	} `json:"Contenido"`
}

type NosisApiConfiguration struct {
	BaseURL  string
	User     string
	Password string
	Group    string
}

type NosisService interface {

	/*
	 * Get user
	 */
	Get(userId string) (*NosisResponse, error)
}

func (r *NosisResponse) Get(varName string) string {
	for _, n := range r.Contenido.Datos.Variables {
		if varName == n.Nombre {
			return n.Valor
		}
	}

	return ""
}
