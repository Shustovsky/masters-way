import {SchemasLabelStatistics} from "src/apiAutogenerated/general";
import {LabelStatistics} from "src/model/businessModel/WayStatistics";

/**
 * Convert {@link SchemasLabelStatistics} to {@link LabelStatistics}
 */
export const labelStatisticsDTOToLabelStatistics = (labelStatisticsDTO: SchemasLabelStatistics): LabelStatistics => {
  const labelStatistics = new LabelStatistics(labelStatisticsDTO);

  return labelStatistics;
};
