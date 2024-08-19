import {Configuration as ChatConfiguration, MessageApi, RoomApi} from "src/apiAutogenerated/chat";
import {
  AuthApi,
  CommentApi,
  CompositeWayApi,
  Configuration as GeneralConfiguration,
  DayReportApi,
  FavoriteUserApi,
  FavoriteUserWayApi,
  FetchParams,
  FromUserMentoringRequestApi,
  GeminiApi,
  HealthApi,
  JobDoneApi,
  JobDoneJobTagApi,
  JobTagApi,
  MentorUserWayApi,
  MetricApi,
  Middleware,
  PlanApi,
  PlanJobTagApi,
  ProblemApi,
  RequestContext,
  ResponseContext,
  ToUserMentoringRequestApi,
  UserApi, UserTagApi,
  WayApi,
  WayCollectionApi,
  WayCollectionWayApi,
  WayTagApi,
} from "src/apiAutogenerated/general";
import {tokenStore} from "src/globalStore/TokenStore";
import {userStore} from "src/globalStore/UserStore";
import {env} from "src/utils/env/env";

/**
 * Access token provider
 */
const accessTokenProvider = () => tokenStore.accessToken ?? "";

const NOT_AUTHORIZED_STATUS = 401;

/**
 * Middleware to help refresh token
 */
const checkAuthMiddleware: Middleware = {

  /**
   * Pre request middle ware
   * @param context
   */
  pre: async (context: RequestContext): Promise<FetchParams | void> => {
    // Retrieve the access token
    const token = tokenStore.accessToken;

    // If a token exists, add it to the Authorization header
    if (token) {
      context.init.headers = {
        ...context.init.headers,
        "Authorization": `Bearer ${token}`,
      };
    }

    return Promise.resolve(context);
  },

  /**
   * Post request middleware
   */
  post: (context: ResponseContext): Promise<Response | void> => {
    if (context.response.status === NOT_AUTHORIZED_STATUS) {
      tokenStore.setAccessToken(null);
      userStore.clearUser();
    }

    return Promise.resolve(undefined);
  },
};

const generalConfiguration = new GeneralConfiguration({
  basePath: env.API_BASE_PATH,
  credentials: "include",
  middleware: [checkAuthMiddleware],
  accessToken: accessTokenProvider,
});

export const healthCheckService = new HealthApi(generalConfiguration);
export const authService = new AuthApi(generalConfiguration);
export const userService = new UserApi(generalConfiguration);
export const wayCollectionService = new WayCollectionApi(generalConfiguration);
export const favoriteUserService = new FavoriteUserApi(generalConfiguration);
export const userTagService = new UserTagApi(generalConfiguration);
export const wayTagService = new WayTagApi(generalConfiguration);
export const wayCollectionWayService = new WayCollectionWayApi(generalConfiguration);
export const wayService = new WayApi(generalConfiguration);
export const toUserMentoringRequestService = new ToUserMentoringRequestApi(generalConfiguration);
export const planJobTagService = new PlanJobTagApi(generalConfiguration);
export const jobDoneJobTagService = new JobDoneJobTagApi(generalConfiguration);
export const jobTagService = new JobTagApi(generalConfiguration);
export const jobDoneService = new JobDoneApi(generalConfiguration);
export const planService = new PlanApi(generalConfiguration);
export const problemService = new ProblemApi(generalConfiguration);
export const commentService = new CommentApi(generalConfiguration);
export const dayReportService = new DayReportApi(generalConfiguration);
export const metricService = new MetricApi(generalConfiguration);
export const favoriteUserWayService = new FavoriteUserWayApi(generalConfiguration);
export const fromUserMentoringRequest = new FromUserMentoringRequestApi(generalConfiguration);
export const mentorUserWay = new MentorUserWayApi(generalConfiguration);
export const compositeWayService = new CompositeWayApi(generalConfiguration);
export const aiService = new GeminiApi(generalConfiguration);

const chatConfiguration = new ChatConfiguration({
  basePath: env.API_CHAT_BASE_PATH,
  credentials: "include",
  middleware: [checkAuthMiddleware],
  accessToken: accessTokenProvider,
});

export const chat = new RoomApi(chatConfiguration);
export const messageService = new MessageApi(chatConfiguration);
