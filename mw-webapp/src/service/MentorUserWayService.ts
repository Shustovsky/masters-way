import {CreateMentorUserWayRequest, DeleteMentorUserWayRequest} from "src/apiAutogenerated/general";
import {mentorUserWay} from "src/service/services";

/**
 * Provides methods to interact with the mentorUserWay
 */
export class MentorUserWayService {

  /**
   * Add mentor to the way
   */
  public static async addMentorUserWay(requestParameters: CreateMentorUserWayRequest): Promise<void> {
    await mentorUserWay.createMentorUserWay(requestParameters);
  }

  /**
   * Delete mentor from the way
   */
  public static async deleteMentorUserWay(requestParameters: DeleteMentorUserWayRequest): Promise<void> {
    await mentorUserWay.deleteMentorUserWay(requestParameters);
  }

}
