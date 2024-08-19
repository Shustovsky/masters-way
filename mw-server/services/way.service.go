package services

import (
	"context"
	"fmt"
	dbb "mwserver/db/sqlc"
	"mwserver/schemas"
	"mwserver/util"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/samber/lo"
)

type GetPopulatedWayByIdParams struct {
	WayUuid              uuid.UUID
	CurrentChildrenDepth int
}

func GetPopulatedWayById(db *dbb.Queries, ctx context.Context, params GetPopulatedWayByIdParams) (schemas.WayPopulatedResponse, error) {
	wayPgUUID := pgtype.UUID{Bytes: params.WayUuid, Valid: true}
	way, err := db.GetWayById(ctx, wayPgUUID)
	if err != nil {
		return schemas.WayPopulatedResponse{}, err
	}

	wayOwner := schemas.UserPlainResponse{
		Uuid:        util.ConvertPgUUIDToUUID(way.OwnerUuid).String(),
		Name:        way.OwnerName,
		Email:       way.OwnerEmail,
		Description: way.OwnerDescription,
		CreatedAt:   way.OwnerCreatedAt.Time.Format(util.DEFAULT_STRING_LAYOUT),
		ImageUrl:    way.OwnerImageUrl,
		IsMentor:    way.OwnerIsMentor,
	}

	jobTagsRaw, _ := db.GetListJobTagsByWayUuid(ctx, wayPgUUID)
	jobTags := lo.Map(jobTagsRaw, func(dbJobTag dbb.JobTag, i int) schemas.JobTagResponse {
		return schemas.JobTagResponse{
			Uuid:        util.ConvertPgUUIDToUUID(dbJobTag.Uuid).String(),
			Name:        dbJobTag.Name,
			Description: dbJobTag.Description,
			Color:       dbJobTag.Color,
		}
	})

	favoriteForUserAmount, _ := db.GetFavoriteForUserUuidsByWayId(ctx, wayPgUUID)
	fromUserMentoringRequestsRaw, _ := db.GetFromUserMentoringRequestWaysByWayId(ctx, wayPgUUID)
	fromUserMentoringRequests := lo.Map(fromUserMentoringRequestsRaw, func(fromUser dbb.User, i int) schemas.UserPlainResponse {
		return schemas.UserPlainResponse{
			Uuid:        util.ConvertPgUUIDToUUID(fromUser.Uuid).String(),
			Name:        fromUser.Name,
			Email:       fromUser.Email,
			Description: fromUser.Description,
			CreatedAt:   fromUser.CreatedAt.Time.Format(util.DEFAULT_STRING_LAYOUT),
			ImageUrl:    fromUser.ImageUrl,
			IsMentor:    fromUser.IsMentor,
		}
	})

	formerMentorsRaw, _ := db.GetFormerMentorUsersByWayId(ctx, wayPgUUID)
	formerMentors := lo.Map(formerMentorsRaw, func(dbFormerMentor dbb.User, i int) schemas.UserPlainResponse {
		return schemas.UserPlainResponse{
			Uuid:        util.ConvertPgUUIDToUUID(dbFormerMentor.Uuid).String(),
			Name:        dbFormerMentor.Name,
			Email:       dbFormerMentor.Email,
			Description: dbFormerMentor.Description,
			CreatedAt:   dbFormerMentor.CreatedAt.Time.Format(util.DEFAULT_STRING_LAYOUT),
			ImageUrl:    dbFormerMentor.ImageUrl,
			IsMentor:    dbFormerMentor.IsMentor,
		}
	})

	mentorsRaw, _ := db.GetMentorUsersByWayId(ctx, wayPgUUID)
	mentors := lo.Map(mentorsRaw, func(dbMentor dbb.User, i int) schemas.UserPlainResponse {
		return schemas.UserPlainResponse{
			Uuid:        util.ConvertPgUUIDToUUID(dbMentor.Uuid).String(),
			Name:        dbMentor.Name,
			Email:       dbMentor.Email,
			Description: dbMentor.Description,
			CreatedAt:   dbMentor.CreatedAt.Time.Format(util.DEFAULT_STRING_LAYOUT),
			ImageUrl:    dbMentor.ImageUrl,
			IsMentor:    dbMentor.IsMentor,
		}
	})

	metricsRaw, _ := db.GetListMetricsByWayUuid(ctx, wayPgUUID)
	metrics := lo.Map(metricsRaw, func(dbMetric dbb.Metric, i int) schemas.MetricResponse {
		return schemas.MetricResponse{
			Uuid:             util.ConvertPgUUIDToUUID(dbMetric.Uuid).String(),
			Description:      dbMetric.Description,
			IsDone:           dbMetric.IsDone,
			DoneDate:         util.MarshalPgTimestamp(dbMetric.DoneDate),
			MetricEstimation: dbMetric.MetricEstimation,
		}
	})

	wayTagsRaw, _ := db.GetListWayTagsByWayId(ctx, wayPgUUID)
	wayTags := lo.Map(wayTagsRaw, func(dbWayTag dbb.WayTag, i int) schemas.WayTagResponse {
		return schemas.WayTagResponse{
			Uuid: util.ConvertPgUUIDToUUID(dbWayTag.Uuid).String(),
			Name: dbWayTag.Name,
		}
	})

	var children []schemas.WayPopulatedResponse
	if params.CurrentChildrenDepth < int(limitMap[MaxCompositeWayDeps][dbb.PricingPlanTypeStarter]) {
		children = lo.Map(way.ChildrenUuids, func(childUuid string, i int) schemas.WayPopulatedResponse {
			args := GetPopulatedWayByIdParams{
				WayUuid:              uuid.MustParse(childUuid),
				CurrentChildrenDepth: params.CurrentChildrenDepth + 1,
			}
			child, _ := GetPopulatedWayById(db, ctx, args)

			return child
		})
	} else {
		children = []schemas.WayPopulatedResponse{}
	}

	response := schemas.WayPopulatedResponse{
		Uuid:                   util.ConvertPgUUIDToUUID(way.Uuid).String(),
		Name:                   way.Name,
		GoalDescription:        way.GoalDescription,
		UpdatedAt:              way.UpdatedAt.Time.Format(util.DEFAULT_STRING_LAYOUT),
		CreatedAt:              way.CreatedAt.Time.Format(util.DEFAULT_STRING_LAYOUT),
		EstimationTime:         way.EstimationTime,
		IsCompleted:            way.IsCompleted,
		IsPrivate:              way.IsPrivate,
		Owner:                  wayOwner,
		Mentors:                mentors,
		FormerMentors:          formerMentors,
		FromUserMentorRequests: fromUserMentoringRequests,
		FavoriteForUsersAmount: int32(favoriteForUserAmount),
		WayTags:                wayTags,
		JobTags:                jobTags,
		Metrics:                metrics,
		CopiedFromWayUuid:      util.MarshalPgUUID(way.CopiedFromWayUuid),
		Children:               children,
	}

	return response, nil
}

func UpdateWayIsCompletedStatus(db *dbb.Queries, ctx context.Context, wayUuid pgtype.UUID) error {

	isCompleted, err := db.IsAllMetricsDone(ctx, wayUuid)

	now := time.Now()
	args := dbb.UpdateWayParams{
		Uuid:        wayUuid,
		IsCompleted: pgtype.Bool{Bool: isCompleted, Valid: true},
		UpdatedAt:   pgtype.Timestamp{Time: now, Valid: true},
	}

	db.UpdateWay(ctx, args)

	return err
}

func GetPlainWayById(db *dbb.Queries, ctx context.Context, wayUuid pgtype.UUID) (schemas.WayPlainResponse, error) {
	way, err := db.GetWayById(ctx, wayUuid)

	mentorsRaw, _ := db.GetMentorUsersByWayId(ctx, way.Uuid)

	mentors := lo.Map(mentorsRaw, func(dbMentor dbb.User, i int) schemas.UserPlainResponse {
		return schemas.UserPlainResponse{
			Uuid:        util.ConvertPgUUIDToUUID(dbMentor.Uuid).String(),
			Name:        dbMentor.Name,
			Email:       dbMentor.Email,
			Description: dbMentor.Description,
			CreatedAt:   dbMentor.CreatedAt.Time.Format(util.DEFAULT_STRING_LAYOUT),
			ImageUrl:    dbMentor.ImageUrl,
			IsMentor:    dbMentor.IsMentor,
		}
	})

	dbOwner, _ := db.GetUserById(ctx, way.OwnerUuid)
	owner := schemas.UserPlainResponse{
		Uuid:        util.ConvertPgUUIDToUUID(dbOwner.Uuid).String(),
		Name:        dbOwner.Name,
		Email:       dbOwner.Email,
		Description: dbOwner.Description,
		CreatedAt:   dbOwner.CreatedAt.Time.Format(util.DEFAULT_STRING_LAYOUT),
		ImageUrl:    dbOwner.ImageUrl,
		IsMentor:    dbOwner.IsMentor,
	}
	dbTags, _ := db.GetListWayTagsByWayId(ctx, way.Uuid)
	wayTags := lo.Map(dbTags, func(dbTag dbb.WayTag, i int) schemas.WayTagResponse {
		return schemas.WayTagResponse{
			Uuid: util.ConvertPgUUIDToUUID(dbTag.Uuid).String(),
			Name: dbTag.Name,
		}
	})
	response := schemas.WayPlainResponse{
		Uuid:              util.ConvertPgUUIDToUUID(way.Uuid).String(),
		Name:              way.Name,
		GoalDescription:   way.GoalDescription,
		UpdatedAt:         way.UpdatedAt.Time.Format(util.DEFAULT_STRING_LAYOUT),
		CreatedAt:         way.CreatedAt.Time.Format(util.DEFAULT_STRING_LAYOUT),
		EstimationTime:    way.EstimationTime,
		IsCompleted:       way.IsCompleted,
		Owner:             owner,
		CopiedFromWayUuid: util.MarshalPgUUID(way.CopiedFromWayUuid),
		IsPrivate:         way.IsPrivate,
		FavoriteForUsers:  int32(way.WayFavoriteForUsers),
		DayReportsAmount:  int32(way.WayDayReportsAmount),
		Mentors:           mentors,
		MetricsDone:       int32(way.WayMetricsDone),
		MetricsTotal:      int32(way.WayMetricsTotal),
		WayTags:           wayTags,
		ChildrenUuids:     way.ChildrenUuids,
	}

	return response, err
}

type GetWayStatisticsTriplePeriodParams struct {
	WayUUIDs       []uuid.UUID
	TotalStartDate time.Time
	EndDate        time.Time
}

func GetWayStatisticsTriplePeriod(db *dbb.Queries, ctx context.Context, params *GetWayStatisticsTriplePeriodParams) (*schemas.WayStatisticsTriplePeriod, error) {
	wayPgUUIDs := lo.Map(params.WayUUIDs, func(wayUUID uuid.UUID, _ int) pgtype.UUID {
		return pgtype.UUID{Bytes: wayUUID, Valid: true}
	})
	endDatePgTimestamp := pgtype.Timestamp{Time: params.EndDate, Valid: true}

	totalTimeStatisticsParams := &GetWayStatisticsParams{
		WayPgUUIDs:           wayPgUUIDs,
		StartDatePgTimestamp: pgtype.Timestamp{Time: params.TotalStartDate, Valid: true},
		EndDatePgTimestamp:   endDatePgTimestamp,
	}
	totalTimeStatistics, err := GetWayStatistics(db, ctx, totalTimeStatisticsParams)
	if err != nil {
		return nil, err
	}

	lastMonthStatisticsParams := &GetWayStatisticsParams{
		WayPgUUIDs:           wayPgUUIDs,
		StartDatePgTimestamp: pgtype.Timestamp{Time: params.EndDate.AddDate(0, -1, 0), Valid: true},
		EndDatePgTimestamp:   endDatePgTimestamp,
	}
	lastMonthStatistics, err := GetWayStatistics(db, ctx, lastMonthStatisticsParams)
	if err != nil {
		return nil, err
	}

	lastWeekStatisticsParams := &GetWayStatisticsParams{
		WayPgUUIDs:           wayPgUUIDs,
		StartDatePgTimestamp: pgtype.Timestamp{Time: params.EndDate.AddDate(0, 0, -6), Valid: true},
		EndDatePgTimestamp:   endDatePgTimestamp,
	}
	lastWeekStatistics, err := GetWayStatistics(db, ctx, lastWeekStatisticsParams)
	if err != nil {
		return nil, err
	}

	return &schemas.WayStatisticsTriplePeriod{
		TotalTime: *totalTimeStatistics,
		LastMonth: *lastMonthStatistics,
		LastWeek:  *lastWeekStatistics,
	}, nil
}

type GetWayStatisticsParams struct {
	WayPgUUIDs           []pgtype.UUID
	StartDatePgTimestamp pgtype.Timestamp
	EndDatePgTimestamp   pgtype.Timestamp
}

func GetWayStatistics(db *dbb.Queries, ctx context.Context, params *GetWayStatisticsParams) (*schemas.WayStatistics, error) {
	timeSpentByDayChart, err := GetTimeSpentByDayChart(db, ctx, params)
	if err != nil {
		return nil, err
	}

	overallInformation, err := GetOverallInformation(db, ctx, params)
	if err != nil {
		return nil, err
	}

	getLabelStatistics, err := GetLabelStatistics(db, ctx, params)
	if err != nil {
		return nil, err
	}

	return &schemas.WayStatistics{
		TimeSpentByDayChart: timeSpentByDayChart,
		LabelStatistics:     *getLabelStatistics,
		OverallInformation:  *overallInformation,
	}, nil
}

func GetTimeSpentByDayChart(db *dbb.Queries, ctx context.Context, params *GetWayStatisticsParams) ([]schemas.TimeSpentByDayPoint, error) {
	timeSpentByDayChartParams := dbb.GetTimeSpentByDayChartParams{
		WayUuids:  params.WayPgUUIDs,
		StartDate: params.StartDatePgTimestamp,
		EndDate:   params.EndDatePgTimestamp,
	}

	timeSpentByDayChartRaw, err := db.GetTimeSpentByDayChart(ctx, timeSpentByDayChartParams)
	if err != nil {
		return nil, err
	}
	fmt.Println("params.StartDatePgTimestamp: ", params.StartDatePgTimestamp)
	fmt.Println("params.EndDatePgTimestamp: ", params.EndDatePgTimestamp)
	fmt.Println("params.WayPgUUIDs: ", params.WayPgUUIDs)

	daysCount := int(params.EndDatePgTimestamp.Time.Sub(params.StartDatePgTimestamp.Time).Hours()/24) + 1
	fmt.Println("daysCount: ", daysCount)
	timeSpentByDayChart := make([]schemas.TimeSpentByDayPoint, 0, daysCount)

	timeSpentByDayMap := lo.SliceToMap(timeSpentByDayChartRaw, func(timeSpentByDay dbb.GetTimeSpentByDayChartRow) (time.Time, int) {
		return timeSpentByDay.PointDate.Time.Truncate(24 * time.Hour), int(timeSpentByDay.PointValue)
	})

	for date := params.StartDatePgTimestamp.Time; !date.After(params.EndDatePgTimestamp.Time); date = date.AddDate(0, 0, 1) {
		truncatedDate := date.Truncate(24 * time.Hour)
		value := timeSpentByDayMap[truncatedDate]

		timeSpentByDayChart = append(timeSpentByDayChart, schemas.TimeSpentByDayPoint{
			Value: value,
			Date:  truncatedDate.Format(util.DEFAULT_STRING_LAYOUT),
		})
	}
	return timeSpentByDayChart, nil
}

func GetOverallInformation(db *dbb.Queries, ctx context.Context, params *GetWayStatisticsParams) (*schemas.OverallInformation, error) {
	overallInformationParams := dbb.GetOverallInformationParams{
		WayUuids:  params.WayPgUUIDs,
		StartDate: params.StartDatePgTimestamp,
		EndDate:   params.EndDatePgTimestamp,
	}

	overallInformationRaw, err := db.GetOverallInformation(ctx, overallInformationParams)
	if err != nil {
		return nil, err
	}

	return &schemas.OverallInformation{
		TotalTime:                 int(overallInformationRaw.TotalTime),
		TotalReports:              int(overallInformationRaw.TotalReports),
		FinishedJobs:              int(overallInformationRaw.FinishedJobs),
		AverageTimePerCalendarDay: int(overallInformationRaw.AverageTimePerCalendarDay),
		AverageTimePerWorkingDay:  int(overallInformationRaw.AverageTimePerWorkingDay),
		AverageJobTime:            int(overallInformationRaw.AverageJobTime),
	}, nil
}

func GetLabelStatistics(db *dbb.Queries, ctx context.Context, params *GetWayStatisticsParams) (*schemas.LabelStatistics, error) {
	labelStatisticsParams := dbb.GetLabelStatisticsParams{
		WayUuids:  params.WayPgUUIDs,
		StartDate: params.StartDatePgTimestamp,
		EndDate:   params.EndDatePgTimestamp,
	}

	labelStatisticsRaw, err := db.GetLabelStatistics(ctx, labelStatisticsParams)
	if err != nil {
		return nil, err
	}

	labelsInfo := lo.Map(labelStatisticsRaw, func(dbLabel dbb.GetLabelStatisticsRow, _ int) schemas.LabelInfo {
		return schemas.LabelInfo{
			Label: schemas.Label{
				ID:          util.ConvertPgUUIDToUUID(dbLabel.LabelUuid).String(),
				Name:        dbLabel.LabelName,
				Color:       dbLabel.LabelColor,
				Description: dbLabel.LabelDescription,
			},
			JobsAmount:           int(dbLabel.JobsAmount),
			JobsAmountPercentage: int(dbLabel.JobsAmountPercentage),
			Time:                 int(dbLabel.JobsTime),
			TimePercentage:       int(dbLabel.JobsTimePercentage),
		}
	})

	return &schemas.LabelStatistics{Labels: labelsInfo}, nil
}

func GetChildrenWayIDs(db *dbb.Queries, ctx context.Context, wayID uuid.UUID, maxDepth int) ([]uuid.UUID, error) {
	wayPgUUID := pgtype.UUID{Bytes: wayID, Valid: true}
	waysRaw, err := GetNestedWayIDs(db, ctx, wayPgUUID, 0, maxDepth)
	if err != nil {
		return nil, err
	}

	ways := lo.Map(waysRaw, func(way pgtype.UUID, _ int) uuid.UUID {
		return util.ConvertPgUUIDToUUID(way)
	})

	return ways, nil
}

func GetNestedWayIDs(db *dbb.Queries, ctx context.Context, parentWayUUID pgtype.UUID, currentDepth int, maxDepth int) ([]pgtype.UUID, error) {
	if currentDepth >= maxDepth {
		return nil, nil
	}

	children, err := db.GetWayChildren(ctx, parentWayUUID)
	if err != nil {
		return nil, err
	}

	for _, wayUUID := range children {
		child, err := GetNestedWayIDs(db, ctx, wayUUID, currentDepth+1, maxDepth)
		if err != nil {
			return nil, err
		}
		children = append(children, child...)
	}

	return children, nil
}
