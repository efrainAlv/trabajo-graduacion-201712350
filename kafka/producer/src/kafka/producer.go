package kafka

import(
	"os"
	"context"
	"fmt"
	"log"
	"errors"
	"time"
	"encoding/json"
	kfk "github.com/segmentio/kafka-go"
	ts "kafka/types"

)

func CreteTopic(topicName string) error{

	_, err := kfk.DialLeader(context.Background(), "tcp", os.Getenv("BROKER"), topicName, 0)
	if err != nil {
		log.Print(err)
		return errors.New("Topic couldn't be created")
	}
	
	fmt.Print("Topic created!")
	return nil
}


func ListTopics(){

	conn, err := kfk.Dial("tcp", os.Getenv("BROKER"))
	if err != nil {
		log.Print(err)
	}
	defer conn.Close()
	
	partitions, err := conn.ReadPartitions()
	if err != nil {
		log.Print(err)
	}
	
	m := map[string]struct{}{}
	
	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}

}


func SendMessage(msg ts.Message, topic string) (error, string){

	conn, err := kfk.DialLeader(context.Background(), "tcp", os.Getenv("BROKER"), topic, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
		return errors.New("Cloudn't connect to topic"), ""
	}

	jsonBytes, error := json.Marshal(msg)
	if error!=nil{
		return errors.New("Failed to write message"), ""
	}else{
		conn.SetWriteDeadline(time.Now().Add(10*time.Second))
		_, err = conn.WriteMessages(
			kfk.Message{Value: jsonBytes},
		)
		if err != nil {
			log.Fatal("Failed to write message:", err)
			return errors.New("Failed to write messages"), ""
		}
	
		return nil, fmt.Sprint("Message sended:", err)
	}
	
}	