import * as Sentry from "@sentry/nextjs";

Sentry.init({
  dsn: "https://<TU-DSN>.ingest.sentry.io/<TU-ID>",
  tracesSampleRate: 0.5,
  integrations: [Sentry.browserTracingIntegration()],
  tracePropagationTargets: [
    "https://nextgo-frontend-latest.onrender.com",  // tu dominio público
    /^\/api/                                         // rutas locales de API
  ],
});
