import {useNavigate} from "react-router-dom";
import clsx from "clsx";
import {observer} from "mobx-react-lite";
import {Accordion, accordionTypes} from "src/component/accordion/Accordion";
import {AdvantageItem} from "src/component/advantageItem/AdvantageItem";
import {Button, ButtonType} from "src/component/button/Button";
import {Footer} from "src/component/footer/Footer";
import {HorizontalContainer} from "src/component/horizontalContainer/HorizontalContainer";
import {Icon, IconSize} from "src/component/icon/Icon";
import {Image} from "src/component/image/Image";
import {ReviewCard} from "src/component/reviewCard/ReviewCard";
import {Slider, SliderItem} from "src/component/slider/Slider";
import {HeadingLevel, Title} from "src/component/title/Title";
import {VerticalContainer} from "src/component/verticalContainer/VerticalContainer";
import {languageStore} from "src/globalStore/LanguageStore";
import {themeStore} from "src/globalStore/ThemeStore";
import {AmountBlock, AmountItem} from "src/land/amountBlock/AmountBlock";
import {HeaderLanding, NavLink} from "src/land/headerLanding/HeaderLanding";
import {ProblemItem} from "src/land/problemItem/ProblemItem";
import {PricePlan} from "src/logic/pricingPage/pricePlan/PricePlan";
import {pricePlans} from "src/logic/pricingPage/pricePlan/pricePlans";
import {pages} from "src/router/pages";
import {LanguageService} from "src/service/LanguageService";
import {renderMarkdown} from "src/utils/markdown/renderMarkdown";
import styles from "src/land/landings/businessLandingPage/BusinessLandingPage.module.scss";

/**
 * Business landing page
 */
export const BusinessLandingPage = observer(() => {
  const {language, setLanguage} = languageStore;
  const {theme, setTheme} = themeStore;

  const navigate = useNavigate();

  const navList: NavLink[] = [
    {
      path: "advantages",
      value: LanguageService.businessLanding.navigation.advantages[language],
    },
    {
      path: "problems",
      value: LanguageService.businessLanding.navigation.solutions[language],
    },
    {
      path: "reviews",
      value: LanguageService.businessLanding.navigation.reviews[language],
    },
    {
      path: "pricing",
      value: LanguageService.businessLanding.navigation.pricing[language],
    },
    {
      path: "questions",
      value: LanguageService.businessLanding.navigation.questions[language],
    },
  ];

  const accordionItems = LanguageService.businessLanding.questions.accordion.map((data) => ({
    trigger: {child: data.question[language]},
    content: {child: renderMarkdown(data.answer[language])},
  }));

  const amountItems: AmountItem[] = [
    {
      id: "00000000-0000-0000-0000-00000000001",
      amount: 30,
      description: LanguageService.businessLanding.amountBlock.mentors[language],
    },
    {
      id: "00000000-0000-0000-0000-00000000002",
      amount: 150,
      description: LanguageService.businessLanding.amountBlock.activeWays[language],
    },
    {
      id: "00000000-0000-0000-0000-00000000003",
      amount: 50,
      description: LanguageService.businessLanding.amountBlock.students[language],
    },
  ];

  const reviewCards: SliderItem[] = [
    {
      id: 1,
      content: (
        <ReviewCard
          gradeAmount={5}
          review={LanguageService.businessLanding.reviews.kirillReview.review[language]}
          reviewerImageUrl="https://drive.google.com/thumbnail?id=1T4XtkEjC7KKgqnox7yE9JttbWdxfzlr4&sz=w1000"
          reviewerName={LanguageService.businessLanding.reviews.kirillReview.mentorName[language]}
          reviewerProfession={LanguageService.businessLanding.reviews.kirillReview.mentorProfession[language]}
        />
      ),
    },
    {
      id: 2,
      content: (
        <ReviewCard
          gradeAmount={5}
          review={LanguageService.businessLanding.reviews.viktarReview.review[language]}
          reviewerImageUrl="https://drive.google.com/thumbnail?id=18oHI9KoiaYvd_UowHyqsJbDLLhmuxPxr&sz=w1000"
          reviewerName={LanguageService.businessLanding.reviews.viktarReview.mentorName[language]}
          reviewerProfession={LanguageService.businessLanding.reviews.viktarReview.mentorProfession[language]}
        />
      ),
    },
    {
      id: 3,
      content: (
        <ReviewCard
          gradeAmount={5}
          review={LanguageService.businessLanding.reviews.viktoryiaReview.review[language]}
          reviewerImageUrl="https://drive.google.com/thumbnail?id=1uEyBTZIon2OFQOOG7pCtkigXYD4YXwc8&sz=w1000"
          reviewerName={LanguageService.businessLanding.reviews.viktoryiaReview.mentorName[language]}
          reviewerProfession={LanguageService.businessLanding.reviews.viktoryiaReview.mentorProfession[language]}
        />
      ),
    },
  ];

  return (
    <>
      <div className={styles.container}>
        <HeaderLanding
          language={language}
          setLanguage={setLanguage}
          theme={theme}
          setTheme={setTheme}
          navList={navList}
        />
        <VerticalContainer className={styles.mainBlock}>
          <VerticalContainer className={styles.titleBlock}>

            {/* Workaround for colorized part of title */}
            <h1 className={styles.title}>
              {`${LanguageService.businessLanding.main.title[language]}`}
              <span className={clsx(styles.title, styles.mw)}>
                {`${LanguageService.businessLanding.main.mastersWay[language]}`}
              </span>
            </h1>
            <p className={styles.titleDescription}>
              {LanguageService.businessLanding.main.description[language]}
            </p>
          </VerticalContainer>
          <Button
            buttonType={ButtonType.PRIMARY}
            value={LanguageService.businessLanding.callToActionButton[language]}
            icon={
              <Icon
                size={IconSize.SMALL}
                name="ArrowRightIcon"
                className={styles.icon}
              />
            }
            onClick={() => navigate(pages.home.getPath({}))}
          />
        </VerticalContainer>
        <VerticalContainer className={styles.imageBlock}>
          <Image
            alt="desktopImage"
            src="https://drive.google.com/thumbnail?id=1Y9Xu3b73odKV7VtuLi3RjWM8H7WKiWoA&sz=w1000"
            className={styles.mainBlockImage}
          />
        </VerticalContainer>

        {/* Can't use VErticalContainer because anchor links don't works */}
        <div
          className={styles.advantagesBlock}
          id="advantages"
        >
          <Title
            className={styles.title}
            level={HeadingLevel.h2}
            text={LanguageService.businessLanding.advantages.title[language]}
            placeholder=""
          />
          <HorizontalContainer className={styles.advantages}>
            <AdvantageItem
              iconName="ClockIcon"
              title={LanguageService.businessLanding.advantages.asynchronousInteraction.title[language]}
              description={LanguageService.businessLanding.advantages.asynchronousInteraction.description[language]}
            />
            <AdvantageItem
              iconName="TrendingUpIcon"
              title={LanguageService.businessLanding.advantages.speedingUpProcesses.title[language]}
              description={LanguageService.businessLanding.advantages.speedingUpProcesses.description[language]}
            />
            <AdvantageItem
              iconName="ActivityIcon"
              title={LanguageService.businessLanding.advantages.monitoringAchievements.title[language]}
              description={LanguageService.businessLanding.advantages.monitoringAchievements.description[language]}
            />
          </HorizontalContainer>
        </div>

        <AmountBlock amountItems={amountItems} />

        <div
          className={styles.problemsBlock}
          id="problems"
        >
          <VerticalContainer className={styles.titleBlock}>
            <Title
              className={styles.title}
              level={HeadingLevel.h2}
              text={LanguageService.businessLanding.solutions.title[language]}
              placeholder=""
            />
            <p className={styles.titleDescription}>
              {LanguageService.businessLanding.solutions.description[language]}
            </p>
          </VerticalContainer>
          <VerticalContainer className={styles.problems}>
            <ProblemItem
              title={LanguageService.businessLanding.solutions.timeSaving.title[language]}
              description={LanguageService.businessLanding.solutions.timeSaving.description[language]}
              imageSrc="https://drive.google.com/thumbnail?id=1TSOIZOIg4uvfWeRrE5SUf44CBLanjTUP&sz=w1000"
            />
            <ProblemItem
              title={LanguageService.businessLanding.solutions.progressMonitoring.title[language]}
              description={LanguageService.businessLanding.solutions.progressMonitoring.description[language]}
              imageSrc="https://drive.google.com/thumbnail?id=1ipbye-Gw_KT81T6KDMFpcfKkyxIO9C3g&sz=w1000"
              isReversed
            />
            <ProblemItem
              title={LanguageService.businessLanding.solutions.focusOnResults.title[language]}
              description={LanguageService.businessLanding.solutions.focusOnResults.description[language]}
              imageSrc="https://drive.google.com/thumbnail?id=1fKbItiABNCIlHNM87qwTH1L0fezqdiab&sz=w1000"
            />
          </VerticalContainer>
        </div>

        <div
          className={styles.reviewBlock}
          id="reviews"
        >
          <div className={styles.titleBlock}>
            <Title
              className={styles.title}
              level={HeadingLevel.h2}
              text={LanguageService.businessLanding.reviews.title[language]}
              placeholder=""
            />
          </div>
          <HorizontalContainer className={styles.reviewList}>
            <Slider
              sliderItems={reviewCards}
              settings={{pagination: false}}
            />
          </HorizontalContainer>
          <Button
            buttonType={ButtonType.PRIMARY}
            value={LanguageService.businessLanding.callToActionButton[language]}
            icon={
              <Icon
                size={IconSize.SMALL}
                name="ArrowRightIcon"
                className={styles.icon}
              />
            }
            onClick={() => navigate(pages.home.getPath({}))}
            className={styles.triesActionButton}
          />
        </div>

        <div
          className={styles.pricingBlock}
          id="pricing"
        >
          <Title
            className={styles.title}
            level={HeadingLevel.h2}
            text={LanguageService.businessLanding.pricing.title[language]}
            placeholder=""
          />
          <HorizontalContainer className={styles.planItems}>
            {
              pricePlans.map((pricePlan) => (
                <PricePlan
                  key={pricePlan.id}
                  pricePlan={pricePlan}
                />
              ))
            }
          </HorizontalContainer>
        </div>

        <div
          className={styles.questionsBlock}
          id="questions"
        >
          <VerticalContainer className={styles.titleBlock}>
            <Title
              className={styles.title}
              level={HeadingLevel.h2}
              text={LanguageService.businessLanding.questions.title[language]}
              placeholder=""
            />
          </VerticalContainer>
          <Accordion
            items={accordionItems}
            type={accordionTypes.MULTIPLE}
            className={styles.accordion}
          />
        </div>

        <HorizontalContainer className={styles.triesBlock}>
          <HorizontalContainer className={styles.triesContainer}>
            <Title
              className={clsx(styles.titleBlock, styles.title, styles.titleYouShouldBlock)}
              level={HeadingLevel.h2}
              text={LanguageService.businessLanding.aboutApp.title[language]}
              placeholder=""
            />
            <VerticalContainer className={styles.triesContentBlock}>
              <p>
                {LanguageService.businessLanding.aboutApp.tryOurApp[language]}
              </p>
              <p>
                {LanguageService.businessLanding.aboutApp.ourAppIs[language]}
              </p>
              <Button
                buttonType={ButtonType.PRIMARY}
                value={LanguageService.businessLanding.callToActionButton[language]}
                icon={
                  <Icon
                    size={IconSize.SMALL}
                    name="ArrowRightIcon"
                    className={styles.icon}
                  />
                }
                onClick={() => navigate(pages.home.getPath({}))}
                className={styles.triesActionButton}
              />
            </VerticalContainer>
          </HorizontalContainer>
        </HorizontalContainer>

      </div>
      <Footer language={language} />
    </>
  );
});
