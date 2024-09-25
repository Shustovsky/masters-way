import {useEffect, useState} from "react";
import {DialogClose} from "@radix-ui/react-dialog";
import {Button} from "src/component/button/Button";
import {HorizontalContainer} from "src/component/horizontalContainer/HorizontalContainer";
import {Loader} from "src/component/loader/Loader";
import {VerticalContainer} from "src/component/verticalContainer/VerticalContainer";
import {AIDAL} from "src/dataAccessLogic/AIDAL";
import {languageStore} from "src/globalStore/LanguageStore";
import {themeStore} from "src/globalStore/ThemeStore";
import {LanguageService} from "src/service/LanguageService";
import styles from "src/logic/wayPage/reportsTable/estimateIssueAiModal/EstimateIssueAiModal.module.scss";

/**
 * Estimate issue Ai modal props
 */
interface EstimateIssueAiModalProps {

  /**
   * Goal description
   */
  goalDescription: string;

  /**
   * Issue description
   */
  issueDescription: string;

}

/**
 * Content for estimate issue Ai modal
 */
export const EstimateIssueAiModal = (props: EstimateIssueAiModalProps) => {
  const [generatedEstimateMessage, setGeneratedEstimateMessage] = useState<string>("");
  const {theme} = themeStore;
  const {language} = languageStore;

  /**
   * Estimate issue by using AI
   */
  const estimateIssueAi = async () => {
    const estimatedMessage = await AIDAL.aiEstimateIssue({
      goal: props.goalDescription,
      issue: props.issueDescription,
    });

    setGeneratedEstimateMessage(estimatedMessage);
  };

  useEffect(() => {
    estimateIssueAi();
  }, []);

  return generatedEstimateMessage.length === 0
    ? (
      <Loader theme={theme} />
    )
    : (
      <VerticalContainer className={styles.estimatedMessageModal}>
        {generatedEstimateMessage}
        <HorizontalContainer>
          <DialogClose asChild>
            <Button
              value={LanguageService.modals.confirmModal.cancelButton[language]}
              onClick={() => {}}
            />
          </DialogClose>
        </HorizontalContainer>
      </VerticalContainer>
    );
};
