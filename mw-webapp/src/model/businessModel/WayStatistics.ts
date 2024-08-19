import {makeAutoObservable} from "mobx";
import {LabelStat, LabelStatParams} from "src/logic/wayPage/wayStatistics/LabelStat";

type LabelStatisticsParams = {

  /**
   * Labels info
   */
  labels: LabelStatParams[];
}

/**
 * Data for overall information statistics block
 */
export class LabelStatistics {

  /**
   * List of labels info
   */
  public labels: LabelStat[];

  constructor(params: LabelStatisticsParams) {
    makeAutoObservable(this);
    this.labels = params.labels.map((labelRaw) => new LabelStat(labelRaw));
  }

}

/**
 * Data for overall information statistics block
 */
class OverallInformation {

  /**
   * Average job time
   */
  public averageJobTime: number;

  /**
   * Average time spent by calendar day
   */
  public averageTimePerCalendarDay: number;

  /**
   * Average time spent by working day
   */
  public averageTimePerWorkingDay: number;

  /**
   * Finished jobs amount
   */
  public finishedJobs: number;

  /**
   * Total records
   */
  public totalReports: number;

  /**
   * Total time
   */
  public totalTime: number;

  constructor(params: OverallInformation) {
    makeAutoObservable(this);
    this.averageJobTime = params.averageJobTime;
    this.averageTimePerCalendarDay = params.averageTimePerCalendarDay;
    this.averageTimePerWorkingDay = params.averageTimePerWorkingDay;
    this.finishedJobs = params.finishedJobs;
    this.totalReports = params.totalReports;
    this.totalTime = params.totalTime;
  }

}

/**
 * Single point for chart "time spent by day"
 */
export class TimeSpentByDayPoint {

  /**
   * Point date
   */
  public date: Date;

  /**
   * Point value
   */
  public value: number;

  constructor(params: TimeSpentByDayPoint) {
    makeAutoObservable(this);
    this.date = new Date(params.date);
    this.value = params.value;
  }

}

/**
 * Statistics related to specific way
 */
export class WayStatistics {

  /**
   * List pf points for "time spent by day" chart
   */
  public timeSpentByDayChart: TimeSpentByDayPoint[];

  /**
   * Info for overall information block
   */
  public overallInformation: OverallInformation;

  /**
   * Label statistics
   */
  public labelStatistics: LabelStatistics;

  constructor(params: WayStatistics) {
    makeAutoObservable(this);
    this.timeSpentByDayChart = params.timeSpentByDayChart.map((pointRaw) => new TimeSpentByDayPoint(pointRaw));
    this.overallInformation = new OverallInformation(params.overallInformation);
    this.labelStatistics = new LabelStatistics(params.labelStatistics);
  }

}

/**
 * Way statistics for 3 base periods
 */
export class WayStatisticsTriple {

  /**
   * Total period
   */
  public totalPeriod: WayStatistics;

  /**
   * Last month
   */
  public lastMonth: WayStatistics;

  /**
   * Last week
   */
  public lastWeek: WayStatistics;

  constructor(params: WayStatisticsTriple) {
    makeAutoObservable(this);
    this.totalPeriod = params.totalPeriod;
    this.lastMonth = params.lastMonth;
    this.lastWeek = params.lastWeek;
  }

}
