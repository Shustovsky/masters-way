import type {PricePlanType} from "src/logic/pricingPage/pricePlan/PricePlan";

export const pricePlans: PricePlanType[] = [
  {
    id: 0,
    theme: "light",
    name: "Intro",
    price: 0,
    period: "free",
    capabilities: {
      ownWays: 10,
      privateWays: 1,
      dayReports: 100,
      userTags: 3,
      mentoringWays: 0,
      customCollections: 0,
      compositeWayDeps: 0,
    },
  },
  {
    id: 1,
    theme: "dark",
    name: "Base",
    price: 5,
    period: "month",
    capabilities: {
      ownWays: 20,
      privateWays: 10,
      dayReports: 200,
      userTags: 5,
      mentoringWays: 20,
      customCollections: 0,
      compositeWayDeps: 0,
    },
  },
  {
    id: 2,
    theme: "light",
    name: "PRO",
    price: 50,
    period: "year",
    capabilities: {
      ownWays: 30,
      privateWays: 10,
      dayReports: 365,
      userTags: 5,
      mentoringWays: 30,
      customCollections: 10,
      compositeWayDeps: 3,
    },
  },
];
