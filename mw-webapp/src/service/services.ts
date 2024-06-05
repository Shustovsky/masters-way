import {
  CommentApi,
  CompositeWayApi,
  Configuration,
  DayReportApi,
  FavoriteUserApi,
  FavoriteUserWayApi,
  FromUserMentoringRequestApi,
  JobDoneApi,
  JobDoneJobTagApi,
  JobTagApi,
  MentorUserWayApi,
  MetricApi,
  PlanApi,
  PlanJobTagApi,
  ProblemApi,
  ProblemJobTagApi,
  ToUserMentoringRequestApi,
  UserApi, UserTagApi,
  WayApi,
  WayCollectionApi,
  WayCollectionWayApi,
  WayTagApi,
} from "src/apiAutogenerated";
import {env} from "src/utils/env/env";

const configuration = new Configuration({basePath: env.API_BASE_PATH});

export const userService = new UserApi(configuration);
export const wayCollectionService = new WayCollectionApi(configuration);
export const favoriteUserService = new FavoriteUserApi(configuration);
export const userTagService = new UserTagApi(configuration);
export const wayTagService = new WayTagApi(configuration);
export const wayCollectionWayService = new WayCollectionWayApi(configuration);
export const wayService = new WayApi(configuration);
export const toUserMentoringRequestService = new ToUserMentoringRequestApi(configuration);
export const problemJobTagService = new ProblemJobTagApi(configuration);
export const planJobTagService = new PlanJobTagApi(configuration);
export const jobDoneJobTagService = new JobDoneJobTagApi(configuration);
export const jobTagService = new JobTagApi(configuration);
export const jobDoneService = new JobDoneApi(configuration);
export const planService = new PlanApi(configuration);
export const problemService = new ProblemApi(configuration);
export const commentService = new CommentApi(configuration);
export const dayReportService = new DayReportApi(configuration);
export const metricService = new MetricApi(configuration);
export const favoriteUserWayService = new FavoriteUserWayApi(configuration);
export const fromUserMentoringRequest = new FromUserMentoringRequestApi(configuration);
export const mentorUserWay = new MentorUserWayApi(configuration);
export const compositeWayService = new CompositeWayApi(configuration);
