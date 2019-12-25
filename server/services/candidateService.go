package services

import (
	"encoding/json"
	"github.com/lo9i/databoss/server/data/repositories"
	"github.com/lo9i/databoss/server/domain"
	"strconv"
	"time"
)

type CandidateServiceImpl struct {
	Repository      repositories.CandidateRepository
	NosisRepository repositories.NosisUserRepository
	NosisService    domain.NosisService
}

func NewCandidateService(repository repositories.CandidateRepository, nRepository repositories.NosisUserRepository,
	nService domain.NosisService) domain.CandidateService {
	return &CandidateServiceImpl{Repository: repository, NosisRepository: nRepository, NosisService: nService}
}

func (s *CandidateServiceImpl) Get(id uint64) *domain.Candidate {
	return s.Repository.Get(id)
}

func (s *CandidateServiceImpl) GetByUserId(userId string) *domain.Candidate {
	candidate := s.Repository.GetByUserId(userId)
	if candidate.Id <= 0 {
		candidate = s.CreateCandidate(userId)
	}
	return candidate
}

func (s *CandidateServiceImpl) CreateCandidate(userId string) *domain.Candidate {
	nosisResponse, err := s.NosisService.Get(userId)
	if err != nil {
		return &domain.Candidate{}
	}
	b, err := json.Marshal(nosisResponse)
	if err != nil {
		return &domain.Candidate{}
	}

	candidate := &domain.Candidate{UserId: userId}
	populateCandidate(nosisResponse, candidate)
	s.Repository.Insert(candidate)
	nosisUser := domain.NosisUser{Raw: string(b), CandidateId: candidate.Id}
	s.NosisRepository.Insert(&nosisUser)

	return candidate
}

func populateCandidate(nosisCandidate *domain.NosisResponse, candidate *domain.Candidate) {
	firstname := nosisCandidate.Get("VI_RazonSocial")
	lastname := nosisCandidate.Get("VI_RazonSocial")
	birthStr := nosisCandidate.Get("VI_FecNacimiento")
	birth, err := time.Parse("2006-01-02", birthStr)
	if err != nil {
	}
	male := nosisCandidate.Get("VI_Sexo") == "M"

	phoneNumber := nosisCandidate.Get("VI_Tel1_Nro")

	street := nosisCandidate.Get("VI_DomAF_Calle")
	numberStr := nosisCandidate.Get("VI_DomAF_Nro")
	var number int
	if numberStr != "" {
		number, _ = strconv.Atoi(numberStr)
	}
	floorStr := nosisCandidate.Get("VI_DomAF_Piso")
	var floor int
	if floorStr != "" {
		floor, _ = strconv.Atoi(floorStr)
	}

	department := nosisCandidate.Get("VI_DomAF_Dto")
	city := nosisCandidate.Get("VI_DomAF_Loc")
	state := nosisCandidate.Get("VI_DomAF_Prov")
	zipCode := nosisCandidate.Get("VI_DomAF_CP")

	candidate.FirstName = firstname
	candidate.LastName = lastname
	candidate.BirthDate = birth
	candidate.Male = male
	candidate.PhoneNumber = phoneNumber
	candidate.Street = street
	candidate.Number = number
	candidate.Floor = floor
	candidate.Department = department
	candidate.City = city
	candidate.State = state
	candidate.ZipCode = zipCode
	//candidate.Jobs = nil
}

func (s *CandidateServiceImpl) Find(where ...interface{}) *[]domain.Candidate {
	return s.Repository.Find(where...)
}
