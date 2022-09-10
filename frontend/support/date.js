export function formatDate(datetime) {
  return datetime.setLocale("id").toFormat("EEEE, dd LLL yyyy");
}
