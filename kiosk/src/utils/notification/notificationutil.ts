import { NotificationProgrammatic as Notification } from "buefy";

const NOTIFICATION_TYPE = {
  DANGER: "is-danger",
  SUCCESS: "is-success",
  WARNING: "is-warning"
};

export default class NotificationUtil {
  constructor() {}

  public displayError(message: string) {
    this.displayNotification(message, NOTIFICATION_TYPE.DANGER);
  }

  public displaySuccess(message: string) {
    this.displayNotification(message, NOTIFICATION_TYPE.SUCCESS);
  }

  private displayNotification(message: string, type: string) {
    Notification.open({
      message: message,
      type: type,
      position: "is-top-right",
      duration: 5000,
      hasIcon: false
    });
  }
}
