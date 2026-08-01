import type { ResolutionResult, ResolutionStatus } from "../types";

export interface CoordinatePoint {
  latitude: number;
  longitude: number;
}

export function parseCoordinateQuery(value: string): CoordinatePoint {
  const parts = value
    .trim()
    .split(/[\s,]+/)
    .filter(Boolean);

  if (parts.length !== 2) {
    throw new Error("Enter two coordinates: latitude and longitude.");
  }

  const first = Number(parts[0]);
  const second = Number(parts[1]);
  if (!Number.isFinite(first) || !Number.isFinite(second)) {
    throw new Error("Enter numeric latitude and longitude values.");
  }

  const firstCannotBeLatitude = Math.abs(first) > 90;
  const latitude = firstCannotBeLatitude ? second : first;
  const longitude = firstCannotBeLatitude ? first : second;

  if (Math.abs(latitude) > 90 || Math.abs(longitude) > 180) {
    throw new Error("Enter a valid latitude and longitude.");
  }

  return { latitude, longitude };
}

export function getResolutionPlace(result: ResolutionResult): string {
  const locality = result.geography.municipality?.name ?? result.geography.county?.name;
  const state = result.geography.stateName;

  if (locality && state && locality !== state) {
    return `${locality}, ${state}`;
  }
  return locality ?? state ?? "Matched location";
}

const resolutionNotices: Record<Exclude<ResolutionStatus, "resolved">, string> = {
  partially_resolved: "Some local adoption details could not be confirmed.",
  local_record_required:
    "Contact the applicable local authority to confirm the records identified below.",
  ambiguous: "The available records support more than one interpretation.",
  conflicting: "The available official records conflict.",
  insufficient_evidence: "No reliable determination is available for this location.",
};

export function getResolutionNotice(status: ResolutionStatus): string | null {
  return status === "resolved" ? null : resolutionNotices[status];
}

export function formatCodeFamily(codeFamily: string | undefined): string {
  if (!codeFamily) {
    return "Building codes";
  }

  const label = codeFamily.replace(/_/g, " ");
  if (codeFamily === "fire_operational") {
    return "Operational fire";
  }
  return label.charAt(0).toUpperCase() + label.slice(1);
}
