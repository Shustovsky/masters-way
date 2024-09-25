import {surveyService} from "src/service/services";

/**
 * SurveyUserIntro params
 */
export interface SurveyUserIntroParams {

  /**
   * Device id
   */
  deviceId: string;

  /**
   * Preferred interface language
   */
  preferredInterfaceLanguage: string;

  /**
   * User's role
   */
  role: string;

  /**
   * User's source
   */
  source: string;

  /**
   * User's experience
   */
  studentExperience: string;

  /**
   * User's goals
   */
  studentGoals: string;

  /**
   * User's reason why registered
   */
  whyRegistered: string;

}

/**
 * Provides methods to interact with the surveys
 */
export class SurveyDAL {

  /**
   * Survey user intro
   */
  public static async surveyUserIntro(params: SurveyUserIntroParams): Promise<void> {
    await surveyService.surveyUserIntro({request: {...params}});
  }

}
