const monthOptions: Array<string> = [
  "January",
  "February",
  "March",
  "April",
  "May",
  "June",
  "July",
  "August",
  "September",
  "October",
  "November",
  "December"
];

export default class DateUtil {
  constructor() {}

  public static formatDateToLongDateDisplay(date: Date): string {
    const year = date.getFullYear();
    const month = date.getMonth();
    const day = date.getDay();

    return day + " " + monthOptions[month] + " " + year;
  }
}
