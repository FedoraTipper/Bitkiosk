
export default class CookieUtil {
  constructor() { }

  public convertHeaderToCookieJson(headerString: string): {} {
    let headerSplit: string[] = headerString.split(";");
    let cookie = headerSplit[0].split("=")[1];
    let path = headerSplit[1].split("=")[1];
    let domain = headerSplit[2].split("=")[1];
    let age = headerSplit[3].split("=")[1];
    let httpOnly = headerSplit[4].trim() === "HttpOnly";
    console.log(headerSplit[5].split("=")[1]);
    console.log(headerSplit)
    let secure = headerSplit[5].trim() === "Secure";

    return {
      value: cookie,
      domain: domain,
      expires: Number(age),
      path: path,
      sameSite: 'lax',
      secure: secure,
      httpOnly: httpOnly
    }

  }
}
