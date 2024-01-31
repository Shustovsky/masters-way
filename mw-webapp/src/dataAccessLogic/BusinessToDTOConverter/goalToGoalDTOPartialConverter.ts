import {deleteUndefinedFields} from "src/dataAccessLogic/BusinessToDTOConverter/deleteUndefinedFields";
import {Goal} from "src/model/businessModel/Goal";
import {GoalDTO, GoalPartialDTOSchema} from "src/model/DTOModel/GoalDTO";
import {PartialWithUuid} from "src/utils/PartialWithUuid";

/**
 * Convert {@link goal} to {@link GoalPartialDTO}
 */
export const goalToGoalDTOPartialConverter = (goal: PartialWithUuid<Goal>): PartialWithUuid<GoalDTO> => {
  const goalPartialDTO: PartialWithUuid<GoalDTO> = {
    uuid: goal.uuid,
    description: goal.description,
    estimationTime: goal.estimationTime,
    studentUuid: goal.student ? goal.student.uuid : undefined,
    metricsStringified: goal.metrics ? goal.metrics.map((goalMetric) => JSON.stringify(goalMetric)) : undefined,
  };

  const preparedGoalPartialDTO = deleteUndefinedFields(goalPartialDTO);

  const validatedGoalPartialDTO = GoalPartialDTOSchema.parse(preparedGoalPartialDTO);

  return validatedGoalPartialDTO;
};