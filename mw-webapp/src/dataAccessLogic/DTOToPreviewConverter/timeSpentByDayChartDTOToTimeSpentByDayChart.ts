import {SchemasTimeSpentByDayPoint} from "src/apiAutogenerated/general";
import {TimeSpentByDayPoint} from "src/model/businessModel/WayStatistics";

/**
 * Convert {@link SchemasTimeSpentByDayPoint} to {@link TimeSpentByDayPoint}
 */
export const timeSpentByDayChartDTOToTimeSpentByDayChart = (
  timeSpentByDayChartDTO: SchemasTimeSpentByDayPoint,
): TimeSpentByDayPoint => {
  const wayStatistics = new TimeSpentByDayPoint({
    ...timeSpentByDayChartDTO,
    date: new Date(timeSpentByDayChartDTO.date),
  });

  return wayStatistics;
};