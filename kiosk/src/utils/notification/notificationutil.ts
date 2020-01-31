import { NotificationProgrammatic as Notification } from "buefy";

export default class NotificationUtil {
  constructor() {}

  public displayError(error: string) {
    Notification.open({
      message: error,
      hasIcon: false,
      position: "is-top-right",
      type: "is-danger",
      duration: 5000
    });
  }
}
