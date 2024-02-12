import { HelloWorldPageRoutesConfig } from "../pages/HelloWorldPage";

import { Navigate } from "react-router-dom";

import FuseUtils from "@fuse/utils";
import FuseLoading from "@fuse/core/FuseLoading";

import { LoginPageConfig } from "src/app/pages/LoginPage/LoginPageConfig";
import { NotFoundPageApp } from "src/app/pages/NotFoundPage/NotFoundPageApp";
import { HomePageApp } from "src/app/pages/HomePage/HomePageApp";
import { ProfilePageConfig } from "src/app/pages/ProfilePage/ProfilePageConfig";
import { SignInPageConfig } from "src/app/pages/SignInPage/SignInPageConfig";
import { SignOutPageConfig } from "src/app/pages/SignOutPage/SignOutPageConfig";

import settingsConfig from "app/configs/settingsConfig";
import UsersAppConfig from "../pages/AdministrationSection/Users/UsersAppConfig";
import ClientsAppConfig from "../pages/CounterpartiesSection/clients/ClientsAppConfig";
import multiCodeAppConfig from "../pages/multiCodeAppConfig";
import LimitsAppConfig from "src/app/pages/PlanningSection/limits/LimitsAppConfig";
import OrderJournalAppConfig from "src/app/pages/PlanningSection/order-journal/OrderJournalAppConfig";
import OrderRequestsAppConfig from "../pages/PlanningSection/order-requests/OrderRequestsAppConfig";
import { OperativeReferenceAppConfig } from "src/app/pages/AnalyticSection/operative-reference/OperativeReferenceConfig";

import { DislocationAppConfig } from "src/app/pages/CarriagesSection/dislocation/DislocationConfig";
import CarriageInfoConfig from "src/app/pages/CarriagesSection/carriage-info/CarriageInfoConfig";
import PaymentJournalAppConfig from "src/app/pages/AccountingSection/payment/PaymentJournalAppConfig";
import OwnersConfig from "src/app/pages/CounterpartiesSection/owners/OwnersConfig";
import { NavigationPageConfig } from "src/app/pages/AdministrationSection/navigation/NavigationPageConfig";
import { MovementJournalConfig } from "src/app/pages/CarriagesSection/movement-journal/MovementJournalConfig";
import NodesAppConfig from "../pages/PlanningSection/nodes/NodesAppConfig";
import { ActsJournalConfig } from "../pages/AccountingSection/acts-journal/ActsJournalConfig";
import { CarriagesConfig } from "../pages/AdministrationSection/carriages/CarriagesConfig";
import { SettingsConfig } from "../pages/AdministrationSection/settings/SettingsConfig";

import { ClientsConfig } from "src/app/pages/CounterpartiesSection/clients-test/ClientsConfig";
import { PlanSettingsConfig } from "src/app/pages/PlanningSection/plan-settings/PlanSettingsConfig";
import { BranchCommandsConfig } from "../pages/PlanningSection/branch-commands/BranchCommandsConfig";

const routeConfigs = [
  LoginPageConfig,
  SignOutPageConfig,
  SignInPageConfig,
  ProfilePageConfig,
  UsersAppConfig,
  ClientsAppConfig,
  multiCodeAppConfig,
  CarriagesConfig,
  DislocationAppConfig,
  CarriageInfoConfig,
  LimitsAppConfig,
  OrderRequestsAppConfig,
  OrderJournalAppConfig,
  PaymentJournalAppConfig,
  OwnersConfig,
  NavigationPageConfig,
  MovementJournalConfig,
  NodesAppConfig,
  ActsJournalConfig,
  OperativeReferenceAppConfig,
  SettingsConfig,
  ClientsConfig,
  PlanSettingsConfig,
  BranchCommandsConfig,
	HelloWorldPageRoutesConfig,
];

const routes = [
  ...FuseUtils.generateRoutesFromConfigs(routeConfigs),
  {
    path: "/",
    element: <Navigate to="home" />,
    auth: settingsConfig.defaultAuth,
  },
  {
    path: "loading",
    element: <FuseLoading />,
  },
  {
    path: "home",
    element: <HomePageApp />,
    auth: settingsConfig.defaultAuth,
  },
  {
    path: "*",
    element: <NotFoundPageApp />,
    auth: settingsConfig.defaultAuth,
  },
];

export default routes;