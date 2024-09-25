import clsx from "clsx";
import {wayDescriptionAccessIds} from "cypress/accessIds/wayDescriptionAccessIds";
import {Avatar, AvatarSize} from "src/component/avatar/Avatar";
import {Button, ButtonType} from "src/component/button/Button";
import {Confirm} from "src/component/confirm/Confirm";
import {HorizontalContainer} from "src/component/horizontalContainer/HorizontalContainer";
import {Icon, IconSize} from "src/component/icon/Icon";
import {Link} from "src/component/link/Link";
import {Separator} from "src/component/separator/Separator";
import {Tooltip} from "src/component/tooltip/Tooltip";
import {VerticalContainer} from "src/component/verticalContainer/VerticalContainer";
import {CompositeWayDAL} from "src/dataAccessLogic/CompositeWayDAL";
import {languageStore} from "src/globalStore/LanguageStore";
import {WayStatus} from "src/logic/waysTable/wayStatus";
import {WayWithoutDayReports} from "src/model/businessModelPreview/WayWithoutDayReports";
import {pages} from "src/router/pages";
import {LanguageService} from "src/service/LanguageService";
import {DateUtils} from "src/utils/DateUtils";
import styles from "src/logic/wayPage/WayChildrenList/WayChildrenList.module.scss";

/**
 * {@link WayChildrenList} props
 */
interface WayChildrenListProps {

    /**
     * Root Way fo view
     */
  way: WayWithoutDayReports;

  /**
   * Child level (root is 0)
   */
  level: number;

    /**
     * Indicates whether the current user is the owner of way
     */
  isOwner: boolean;

  /**
   * Callback update day reports on trigger
   */
  onUpdateDayReports?: (userUuid: string, wayUuid: string) => void;

}

const LEVEL_INCREMENT = 1;

/**
 * Recursively sorts an array of WayWithoutDayReports objects by their lastUpdate date
 */
const sortWayChildren = (children: WayWithoutDayReports[]): WayWithoutDayReports[] => {
  const sortedChildrenItem = [...children].sort((a, b) => DateUtils.compareDatesDESC(a.lastUpdate, b.lastUpdate));

  return sortedChildrenItem;
};

/**
 * Item for way children list
 */
export const WayChildrenList = (props: WayChildrenListProps) => {
  const {language} = languageStore;

  /**
   * ChildrenItem
   */
  const renderChildrenItem = (child: WayWithoutDayReports) => {
    const isAbandoned = child.status === WayStatus.abandoned;

    return (
      <div key={child.uuid}>
        <VerticalContainer>
          <HorizontalContainer className={styles.childWay}>
            <HorizontalContainer>
              {"*".repeat(props.level)}
              <Avatar
                alt={child.owner.name}
                src={child.owner.imageUrl}
                size={AvatarSize.SMALL}
              />
              <VerticalContainer className={clsx(isAbandoned && styles.abandonedWay)}>
                <Link
                  path={pages.way.getPath({uuid: child.uuid})}
                  className={styles.participantWay}
                  dataCy={wayDescriptionAccessIds.peopleBlock.childLink(child.name)}
                >
                  {child.name}
                </Link>
                <Link
                  path={pages.user.getPath({uuid: child.owner.uuid})}
                  className={styles.participantWay}
                  dataCy={wayDescriptionAccessIds.peopleBlock.childLink(child.owner.name)}
                >
                  {child.owner.name}
                </Link>
                {child.status}
              </VerticalContainer>
            </HorizontalContainer>

            {props.isOwner && (
              <Confirm
                trigger={
                  <Tooltip content={LanguageService.way.peopleBlock.deleteFromComposite[language]}>
                    <Button
                      onClick={() => {}}
                      buttonType={ButtonType.ICON_BUTTON_WITHOUT_BORDER}
                      value={
                        <Icon
                          size={IconSize.SMALL}
                          name="RemoveIcon"
                        />}
                      dataCy={wayDescriptionAccessIds.peopleBlock.deleteFromCompositeWayButton(child.name)}
                    />
                  </Tooltip>
                }
                content={<p>
                  {LanguageService.way.peopleBlock.deleteWayFromCompositeModalContent[language]
                    .replace("$participant", child.name)}
                </p>}
                onOk={async () => {
                  await CompositeWayDAL.deleteWayFromComposite({childWayUuid: child.uuid, parentWayUuid: props.way.uuid});
                  props.onUpdateDayReports && props.onUpdateDayReports(child.owner.uuid, child.uuid);
                  // Props.way.deleteChildWay(child.uuid);
                }}
                okText={LanguageService.modals.confirmModal.deleteButton[language]}
                cancelText={LanguageService.modals.confirmModal.cancelButton[language]}
                cy={{onOk: wayDescriptionAccessIds.peopleBlock.dialogContent.deleteButton}}
              />
            )}
          </HorizontalContainer>
          <Separator />

          <WayChildrenList
            way={child}
            level={props.level + LEVEL_INCREMENT}
            isOwner={false}
          />
        </VerticalContainer>
      </div>
    );
  };

  const sortedChildren = sortWayChildren(props.way.children);

  const childrenList = sortedChildren.map(renderChildrenItem);

  return childrenList;
};
