import {makeAutoObservable} from "mobx";
import {Room} from "src/model/businessModel/Chat";

/**
 * All chatRoom block-related methods
 */
export class ActiveChatStore {

  /**
   * Active chat
   */
  public activeChat: Room;

  /**
   * If true then active chat is hidden and only chat list is shown
   */
  public isChatHidden: boolean = true;

  /**
   * Message's value
   */
  public message: string = "";

  constructor(activeChat: Room) {
    makeAutoObservable(this);
    this.activeChat = activeChat;
  }

  /**
   * Set is chat hidden on mobile
   */
  public setIsChatHiddenOnMobile = (isChatHiddenOnMobile: boolean) => {
    this.isChatHidden = isChatHiddenOnMobile;
  };

  /**
   * Set message to send
   */
  public setMessage = (message: string) => {
    this.message = message;
  };

}

