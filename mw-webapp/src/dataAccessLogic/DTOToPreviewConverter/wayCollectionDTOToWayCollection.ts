import {SchemasWayCollectionPlainResponse} from "src/apiAutogenerated";
import {WayCollection} from "src/model/businessModel/User";

/**
 * Convert {@link UserDTO} to {@link User}
 */
export const WayCollectionDTOToWayCollection = (wayCollectionDTO: SchemasWayCollectionPlainResponse): WayCollection => {
  const wayCollection: WayCollection = {
    ...wayCollectionDTO,
    uuid: wayCollectionDTO.uuid,
    name: wayCollectionDTO.name,
    ways: [],
    createdAt: new Date(),
    ownerUuid: "",
    updatedAt: new Date(),
    //TODO: check incoming type
    type: "",
  };

  return wayCollection;
};