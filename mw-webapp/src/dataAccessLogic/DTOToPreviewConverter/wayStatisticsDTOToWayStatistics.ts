import {SchemasWayStatistics} from "src/apiAutogenerated/general";
import {labelStatisticsDTOToLabelStatistics} from "src/dataAccessLogic/DTOToPreviewConverter/labelStatisticsDTOToLabelStatistics";
import {timeSpentByDayChartDTOToTimeSpentByDayChart} from
  "src/dataAccessLogic/DTOToPreviewConverter/timeSpentByDayChartDTOToTimeSpentByDayChart";
import {WayStatistics} from "src/model/businessModel/WayStatistics";

/**
 * Convert {@link SchemasWayStatistics} to {@link WayStatistics}
 */
export const wayStatisticsDTOToWayStatistics = (wayStatisticsDTO: SchemasWayStatistics): WayStatistics => {
  const wayStatistics = new WayStatistics({
    ...wayStatisticsDTO,
    timeSpentByDayChart: wayStatisticsDTO.timeSpentByDayChart.map(timeSpentByDayChartDTOToTimeSpentByDayChart),
    labelStatistics: labelStatisticsDTOToLabelStatistics(wayStatisticsDTO.labelStatistics),
  });

  return wayStatistics;
};
