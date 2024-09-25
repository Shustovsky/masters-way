import {SchemasCompositeDayReportPopulatedResponse} from "src/apiAutogenerated/general";
import {commentDTOToComment} from "src/dataAccessLogic/DTOToPreviewConverter/commentDTOToComment";
import {jobDoneDTOToJobDone} from "src/dataAccessLogic/DTOToPreviewConverter/jobDoneDTOToJobDone";
import {planDTOToPlan} from "src/dataAccessLogic/DTOToPreviewConverter/planDTOToPlan";
import {problemDTOToProblem} from "src/dataAccessLogic/DTOToPreviewConverter/problemDTOToProblem";
import {DayReport} from "src/model/businessModel/DayReport";

/**
 * Convert {@link DayReportDTO} to {@link DayReport}
 */
export const dayReportDTOToDayReport = (dayReportDTO: SchemasCompositeDayReportPopulatedResponse): DayReport => {
  const plans = dayReportDTO.plans.map(planDTOToPlan);
  const jobsDone = dayReportDTO.jobsDone.map(jobDoneDTOToJobDone);
  const problems = dayReportDTO.problems.map(problemDTOToProblem);
  const comments = dayReportDTO.comments.map(commentDTOToComment);

  return new DayReport({
    ...dayReportDTO,
    createdAt: new Date(dayReportDTO.createdAt),
    updatedAt: new Date(dayReportDTO.updatedAt),
    jobsDone,
    plans,
    problems,
    comments,
  });
};
