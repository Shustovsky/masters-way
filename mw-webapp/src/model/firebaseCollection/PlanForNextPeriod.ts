/**
 * Plan for next period (day, month, etc.)
 */
export class PlanForNextPeriod {

  /**
   * Plan's UUID
   */
  public uuid: string;

  /**
   * Task that should be done in next period
   */
  public job: string;

  /**
   * Enum @Time.unit (minute, hour, day, etc.)
   */
  public timeUnit: string;

  /**
   * Estimation time for complete job
   */
  public estimationTime: number;

  constructor(planForNextPeriodData: PlanForNextPeriod) {
    this.uuid = planForNextPeriodData.uuid;
    this.job = planForNextPeriodData.job;
    this.timeUnit = planForNextPeriodData.timeUnit;
    this.estimationTime = planForNextPeriodData.estimationTime;
  }

}
