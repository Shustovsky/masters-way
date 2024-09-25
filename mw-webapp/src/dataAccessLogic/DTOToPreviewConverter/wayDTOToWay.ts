import {SchemasWayPopulatedResponse} from "src/apiAutogenerated/general";
import {DayReportsParams} from "src/dataAccessLogic/DayReportDAL";
import {wayWithoutDayReportsDTOToWayWithoutDayReports} from
  "src/dataAccessLogic/DTOToPreviewConverter/wayWithoutDayreportsDTOToWayWithoutDayReports";
import {getWayStatus, WayStatus} from "src/logic/waysTable/wayStatus";
import {Label} from "src/model/businessModel/Label";
import {Metric} from "src/model/businessModel/Metric";
import {UserPlain} from "src/model/businessModel/User";
import {Way} from "src/model/businessModel/Way";
import {arrayToHashMap} from "src/utils/arrayToHashMap";

/**
 * Convert {@link SchemasWayPopulatedResponse} to {@link Way}
 */
export const wayDTOToWay = (wayDTO: SchemasWayPopulatedResponse, dayReports: DayReportsParams): Way => {
  const status = getWayStatus({
    status: wayDTO.isCompleted ? WayStatus.completed : null,
    lastUpdate: new Date(wayDTO.updatedAt),
  });

  const owner = new UserPlain({
    ...wayDTO.owner,
    createdAt: new Date(wayDTO.owner.createdAt),
  });

  const mentors = wayDTO.mentors.map((mentor) => new UserPlain({
    ...mentor,
    createdAt: new Date(mentor.createdAt),
  }));

  const formerMentors = wayDTO.formerMentors.map((formerMentor) => new UserPlain({
    ...formerMentor,
    createdAt: new Date(formerMentor.createdAt),
  }));

  return new Way({
    ...wayDTO,
    metrics: wayDTO.metrics.map((metric) => {
      return new Metric({
        ...metric,
        doneDate: metric.doneDate ? new Date(metric.doneDate) : null,
      });
    }),
    status,
    favoriteForUsersAmount: wayDTO.favoriteForUsersAmount,
    mentorRequests: wayDTO.mentorRequests.map((mentorRequest) => new UserPlain({
      ...mentorRequest,
      createdAt: new Date(mentorRequest.createdAt),
    })),
    mentors: arrayToHashMap({keyField: "uuid", list: mentors}),
    formerMentors: arrayToHashMap({keyField: "uuid", list: formerMentors}),
    dayReports: dayReports.dayReports,
    dayReportsAmount: dayReports.size,
    createdAt: new Date(wayDTO.createdAt),
    lastUpdate: new Date(wayDTO.updatedAt),
    owner,
    children: wayDTO.children.map(wayWithoutDayReportsDTOToWayWithoutDayReports),
    jobTags: wayDTO.jobTags.map(label => new Label(label)),
  });
};
