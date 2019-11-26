package org.coolbeevip;

import javax.ws.rs.Consumes;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;

@Path("/")
public class GreetingResource {

  @POST
  @Produces(MediaType.APPLICATION_JSON)
  @Consumes(MediaType.APPLICATION_JSON)
  public Message hello(Message msg) {
    Message replyMsg = new Message();
    if(msg == null){
      replyMsg.setText("Hi,I'm OpenFaaS. Nothing to say?");
    }else{
      replyMsg.setText("Hi,I'm OpenFaaS. I have received your message '"+msg.getText()+"'");
    }

    return replyMsg;
  }
}