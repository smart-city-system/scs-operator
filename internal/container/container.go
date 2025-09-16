package container

import (
	alarm_repository "scs-operator/internal/app/alarm/repository"
	alarm_service "scs-operator/internal/app/alarm/service"
	guard_premise_repository "scs-operator/internal/app/guard/repository"
	guard_repository "scs-operator/internal/app/guard/repository"
	guard_service "scs-operator/internal/app/guard/service"
	guidance_step_repository "scs-operator/internal/app/guidance-step/repository"
	guidance_step_service "scs-operator/internal/app/guidance-step/service"
	guidance_template_repository "scs-operator/internal/app/guidance-template/repository"
	guidance_template_service "scs-operator/internal/app/guidance-template/service"
	incident_repository "scs-operator/internal/app/incident/repository"
	incident_service "scs-operator/internal/app/incident/service"
	premise_repository "scs-operator/internal/app/premise/repository"
	premise_service "scs-operator/internal/app/premise/service"
	user_repository "scs-operator/internal/app/user/repository"
	kafka_client "scs-operator/pkg/kafka"

	"gorm.io/gorm"
)

// Container holds all the application dependencies
type Container struct {
	// Repositories
	AlarmRepo                *alarm_repository.AlarmRepository
	PremiseRepo              *premise_repository.PremiseRepository
	IncidentRepo             *incident_repository.IncidentRepository
	IncidentGuidanceRepo     *incident_repository.IncidentGuidanceRepository
	IncidentGuidanceStepRepo *incident_repository.IncidentGuidanceStepRepository
	UserRepo                 *user_repository.UserRepository
	GuidanceTemplateRepo     *guidance_template_repository.GuidanceTemplateRepository
	GuidanceStepRepo         *guidance_step_repository.GuidanceStepRepository
	GuardRepo                *guard_repository.GuardRepository
	GuardPremiseRepo         *guard_premise_repository.GuardPremiseRepository

	// Services
	AlarmService            *alarm_service.Service
	PremiseService          *premise_service.Service
	IncidentService         *incident_service.Service
	GuidanceTemplateService *guidance_template_service.Service
	GuidanceStepService     *guidance_step_service.Service
	GuardService            *guard_service.Service
}

// NewContainer creates a new dependency container with all repositories and services
func NewContainer(db *gorm.DB, producer *kafka_client.Producer) *Container {
	// Initialize repositories
	alarmRepo := alarm_repository.NewAlarmRepository(db)
	premiseRepo := premise_repository.NewPremiseRepository(db)
	premiseUsersRepo := premise_repository.NewPremiseUsersRepository(db)
	incidentRepo := incident_repository.NewIncidentRepository(db)
	incidentGuidanceRepo := incident_repository.NewIncidentGuidanceRepository(db)
	incidentGuidanceStepRepo := incident_repository.NewIncidentGuidanceStepRepository(db)
	userRepo := user_repository.NewUserRepository(db)
	guidanceTemplateRepo := guidance_template_repository.NewGuidanceTemplateRepository(db)
	guidanceStepRepo := guidance_step_repository.NewGuidanceStepRepository(db)
	guardPremiseRepo := guard_premise_repository.NewGuardPremiseRepository(db)
	guardRepo := guard_repository.NewGuardRepository(db)

	// Initialize services
	alarmService := alarm_service.NewAlarmService(*alarmRepo, *premiseRepo, *producer)
	premiseService := premise_service.NewPremiseService(*premiseRepo, *premiseUsersRepo)
	incidentService := incident_service.NewIncidentService(*incidentRepo, *incidentGuidanceRepo, *userRepo, *guidanceTemplateRepo, *incidentGuidanceStepRepo, *producer)
	guidanceTemplateService := guidance_template_service.NewGuidanceTemplateService(*guidanceTemplateRepo, *guidanceStepRepo)
	guidanceStepService := guidance_step_service.NewGuidanceStepService(*guidanceStepRepo)
	guardService := guard_service.NewGuardService(*guardRepo, *guardPremiseRepo)

	return &Container{
		// Repositories
		AlarmRepo:                alarmRepo,
		PremiseRepo:              premiseRepo,
		IncidentRepo:             incidentRepo,
		IncidentGuidanceRepo:     incidentGuidanceRepo,
		IncidentGuidanceStepRepo: incidentGuidanceStepRepo,
		UserRepo:                 userRepo,
		GuidanceTemplateRepo:     guidanceTemplateRepo,
		GuidanceStepRepo:         guidanceStepRepo,
		GuardRepo:                guardRepo,
		GuardPremiseRepo:         guardPremiseRepo,

		// Services
		AlarmService:            alarmService,
		PremiseService:          premiseService,
		IncidentService:         incidentService,
		GuidanceTemplateService: guidanceTemplateService,
		GuidanceStepService:     guidanceStepService,
		GuardService:            guardService,
	}
}
