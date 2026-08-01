import type { RefreshStatus } from "../types";

interface StatusBannerProps {
  refreshStatus: RefreshStatus | null;
  apiBaseUrl: string;
  error?: string | null;
}

export function StatusBanner({
  refreshStatus,
  apiBaseUrl,
  error = null,
}: StatusBannerProps): JSX.Element {
  if (!refreshStatus) {
    return (
      <div className="status-banner status-banner--warning">
        <div>
          <strong>{error ?? "Loading cached boundary data."}</strong>
          <p>Waiting for the runtime snapshot to load.</p>
        </div>
        <div className="status-banner__aside">
          <span>API</span>
          <strong>{apiBaseUrl}</strong>
        </div>
      </div>
    );
  }

  return (
    <div className={`status-banner status-banner--${refreshStatus.status}`}>
      <div>
        <strong>{error ?? refreshStatus.message}</strong>
        <p>
          Latest success: {refreshStatus.latestSuccessfulRefresh ?? "none"} | Next scheduled:{" "}
          {refreshStatus.nextScheduledRefresh}
        </p>
      </div>
      <div className="status-banner__aside">
        <span>API</span>
        <strong>{apiBaseUrl}</strong>
      </div>
    </div>
  );
}
