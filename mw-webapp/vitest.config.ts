import {configDefaults, defineConfig} from "vitest/config";
import react from "@vitejs/plugin-react";
import viteTsconfigPaths from "vite-tsconfig-paths";

export default defineConfig({
    plugins: [
        react(),
        viteTsconfigPaths(),
      ],
    test: {
      globals: true,
      environment: "jsdom",
      setupFiles: "./src/setupTests.ts",
      css: true,
      exclude: [...configDefaults.exclude],
      coverage: {
        exclude: [
          "build/",
          "cypress/",
          "*.config.{ts,js}",
          ".storybook/",
          "src/**/*.{test,cy,stories}.tsx",
          "storybook-static/",
          "apiAutogenerated/",
          "assets/",
        ],
      },
    },
  }
);