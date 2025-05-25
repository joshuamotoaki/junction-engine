import js from "@eslint/js";
import typescript from "@typescript-eslint/eslint-plugin";
import typescriptParser from "@typescript-eslint/parser";
import prettier from "eslint-plugin-prettier";
import prettierConfig from "eslint-config-prettier";

export default [
    {
        ignores: ["dist/**", "*.config.js"],
    },
    js.configs.recommended,
    {
        files: ["**/*.{js,ts}"],
        plugins: {
            "@typescript-eslint": typescript,
            prettier,
        },
        languageOptions: {
            parser: typescriptParser,
            ecmaVersion: 2022,
            sourceType: "module",
            parserOptions: {
                project: "./tsconfig.json",
            },
            globals: {
                console: "readonly",
                process: "readonly",
                Buffer: "readonly",
                __dirname: "readonly",
                __filename: "readonly",
                // Bun globals
                Bun: "readonly",
            },
        },
        rules: {
            ...typescript.configs.recommended.rules,
            "@typescript-eslint/no-unused-vars": "warn",
            "@typescript-eslint/explicit-function-return-type": "off",
            "prettier/prettier": "error",
            // API-specific rules
            "no-console": "off", // Allow console logs in API
        },
    },
    prettierConfig,
];
