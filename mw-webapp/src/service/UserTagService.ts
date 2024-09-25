import {
  CreateUserTagRequest,
  DeleteUserTagRequest,
  SchemasUserTagResponse,
} from "src/apiAutogenerated/general";
import {userTagService} from "src/service/services";

/**
 * Provides methods to interact with the user tags
 */
export class UserTagService {

  /**
   * Create user tag
   */
  public static async createUserTag(requestParameters: CreateUserTagRequest): Promise<SchemasUserTagResponse> {
    const userTag = await userTagService.createUserTag(requestParameters);

    return userTag;
  }

  /**
   * Delete user tag by UUID
   */
  public static async deleteUserTag(requestParameters: DeleteUserTagRequest): Promise<void> {
    await userTagService.deleteUserTag(requestParameters);
  }

}
