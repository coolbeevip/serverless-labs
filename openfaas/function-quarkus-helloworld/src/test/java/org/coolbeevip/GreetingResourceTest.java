package org.coolbeevip;

import io.quarkus.test.junit.QuarkusTest;
import javax.ws.rs.core.MediaType;
import org.junit.jupiter.api.Test;

import static io.restassured.RestAssured.given;
import static org.hamcrest.CoreMatchers.startsWith;

/**
 * @author zhanglei
 */
@QuarkusTest
public class GreetingResourceTest {

    @Test
    public void testHelloEndpoint() {
        given()
            .body("{\"text\": \"ping\"}")
            .header("Content-Type", MediaType.APPLICATION_JSON)
            .when().post("/")
            .then()
            .statusCode(200)
            .body(startsWith("{\"text\":\"Hi,I'm OpenFaaS. I have received your message 'ping'\"}"));
    }

}