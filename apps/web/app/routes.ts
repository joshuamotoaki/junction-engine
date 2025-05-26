import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
    index("routes/landing.tsx"),
    route("tos", "routes/tos.tsx"),
    route("privacy", "routes/privacy.tsx"),
    route("snatch", "routes/snatch.tsx"),
    route("recalplus", "routes/recal.tsx"),
    route("coursegenie", "routes/genie.tsx"),
    route("courses", "routes/courses.tsx"),
] satisfies RouteConfig;
