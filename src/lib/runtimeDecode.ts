export class PayloadDecodeError extends Error {
  readonly path: string;

  constructor(path: string, message: string) {
    super(`${path}: ${message}`);
    this.name = "PayloadDecodeError";
    this.path = path;
  }
}

export function record(value: unknown, path = "payload"): Record<string, unknown> {
  if (typeof value !== "object" || value === null || Array.isArray(value)) {
    throw new PayloadDecodeError(path, "expected an object");
  }
  return value as Record<string, unknown>;
}

export function stringValue(value: unknown, path: string): string {
  if (typeof value !== "string") {
    throw new PayloadDecodeError(path, "expected a string");
  }
  return value;
}

export function nonEmptyString(value: unknown, path: string): string {
  const result = stringValue(value, path);
  if (result.trim() === "") {
    throw new PayloadDecodeError(path, "expected a non-empty string");
  }
  return result;
}

export function optionalString(value: unknown, path: string): string | undefined {
  return value === undefined || value === null ? undefined : stringValue(value, path);
}

export function booleanValue(value: unknown, path: string): boolean {
  if (typeof value !== "boolean") {
    throw new PayloadDecodeError(path, "expected a boolean");
  }
  return value;
}

export function finiteNumber(value: unknown, path: string): number {
  if (typeof value !== "number" || !Number.isFinite(value)) {
    throw new PayloadDecodeError(path, "expected a finite number");
  }
  return value;
}

export function longitude(value: unknown, path: string): number {
  const result = finiteNumber(value, path);
  if (result < -180 || result > 180) {
    throw new PayloadDecodeError(path, "expected longitude from -180 through 180");
  }
  return result;
}

export function latitude(value: unknown, path: string): number {
  const result = finiteNumber(value, path);
  if (result < -90 || result > 90) {
    throw new PayloadDecodeError(path, "expected latitude from -90 through 90");
  }
  return result;
}

export function enumValue<const T extends readonly string[]>(
  value: unknown,
  allowed: T,
  path: string,
): T[number] {
  if (typeof value !== "string" || !allowed.includes(value)) {
    throw new PayloadDecodeError(path, `expected one of ${allowed.join(", ")}`);
  }
  return value as T[number];
}

export function arrayValue<T>(
  value: unknown,
  path: string,
  decode: (item: unknown, path: string) => T,
): T[] {
  if (!Array.isArray(value)) {
    throw new PayloadDecodeError(path, "expected an array");
  }
  return value.map((item, index) => decode(item, `${path}[${index}]`));
}

export function optionalArray<T>(
  value: unknown,
  path: string,
  decode: (item: unknown, path: string) => T,
): T[] {
  return value === undefined || value === null ? [] : arrayValue(value, path, decode);
}

export function isoDateTime(value: unknown, path: string): string {
  const result = nonEmptyString(value, path);
  if (!/^\d{4}-\d{2}-\d{2}T/.test(result) || Number.isNaN(Date.parse(result))) {
    throw new PayloadDecodeError(path, "expected an ISO date-time");
  }
  return result;
}

export function isoDate(value: unknown, path: string): string {
  const result = nonEmptyString(value, path);
  if (!/^\d{4}-\d{2}-\d{2}$/.test(result) || Number.isNaN(Date.parse(`${result}T00:00:00Z`))) {
    throw new PayloadDecodeError(path, "expected an ISO date");
  }
  return result;
}

export function httpUrl(value: unknown, path: string): string {
  const result = nonEmptyString(value, path);
  let parsed: URL;
  try {
    parsed = new URL(result);
  } catch {
    throw new PayloadDecodeError(path, "expected an absolute URL");
  }
  if (parsed.protocol !== "http:" && parsed.protocol !== "https:") {
    throw new PayloadDecodeError(path, "expected an HTTP or HTTPS URL");
  }
  return result;
}

export function stringArray(value: unknown, path: string): string[] {
  return arrayValue(value, path, stringValue);
}
