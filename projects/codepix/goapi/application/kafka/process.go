package kafka

import (
	ckafka "github.cm/confluentine/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
	appmodel "github.com/Math-O5/Go/projects/codepix/goapi/application/model/"
)

type KafkaProcessor struct {
	Database *gorm.db
	Producer *ckafka.Producer
	DeliveryChan chan ckafka.Event
}

func NewKafKaProcessor(database *gorm.DB, producer *ckafka.Producer, deliveryChan chan ckafka.Event) *KafkaProcessor {
	return &KafkaProcessor{
		Database: database,
		Producer: producer,
		DeliveryChan: deliveryChan,
	}
}

func (k *KafkaProcessor) Consume() {

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers":os.Getenv("kafkaBootstrapServers"),
		"group.id":os.Getenv("kafkaConsumerGroupId"),
		"auto.offset.reset":"earliest",
	}

	c, err := ckafka.NewConsumer(configMap)

	if err != nil {
		panic(err)
	}

	topics :=  []string{os.Getenv("kafkaTransactionTopic"), os.Getenv("kafkaTransactionConfirmationTopic")}

	c.SubscribeTopics(topics, nil)

	fmt.Println("kafka consumer has been started")

	for {
		msg, err := c.ReadMessage(-1)

		if err != nil {
			fmt.Println(string(msg.Value))
			k.processMessage(msg)
		}
	}
}

func (k *KafkaProcessor) processMessage(msg *ckafka.Message) {
	transactionTopic := "transactions"
	transactionConfirmationTopic := "transaction_confirmation"

	switch topic := *msg.TopicPartition.Topic; topic {
	case transactionTopic:
		k.processTransaction(msg)
	case transactionConfirmationTopic:
		k.transactionConfirmationTopic(msg)
	default:
		fmt.Pritln("not a vaid topic", string(msg.Value))
	}

}


func (k *KafkaProcessor) processTransaction(msg *ckakfa.Message) error {
	transaction := appmodel.NewTransaction()
	err := transaction.PaserJson(msg.Value)
	
	if err != nil {
		return err
	}

	tranactionUseCase := factory.TransactionUseCaseFctory(k.database)

	createdTransaction, err =: tranactionUseCase.Register(
		transaction.AccountId,
		transaction.Amount,
		transaction.PixKeyTo,
		transaction.PixKeyKindTo,
		transaction.Description
	)

	topic := "bank"+ createdTransaction.PixKey.Account.Bank.Code
	transaction.ID = createdTransaction.ID
	transaction.Status = model.TransactionPending
	transactionJson, err := transaction.ToJson() 

	if err != nil {
		fmt.Println("Erro registering transaction...")
		return err
	}

	err := Publish(string(transactionJson), topic, k.Producer, k.DeliveryChan)

	if err != nil {
		fmt.Println("Erro publishing transaction...")
		return err
	}

	return nil
}

func (k *KafkaProcessor) processTransactionConfirmation(msg *ckakfa.Message) error {
	transaction := appmodel.NewTransaction()
	err := transaction.PaserJson(msg.Value)
		
	if err != nil {
		return err
	}

	tranactionUseCase := factory.TransactionUseCaseFctory(k.database)

	if tranactionUseCase.Status == model.TransactionConfirmed {
		err = k.confirmTransaction(transaction, transactionUseCase)

		if err != nil {
			return err
		}

	} else if tranactionUseCase.Status == model.TransactionCompleted {
		_, err := transactionUseCase.Complete(transaction.ID)

		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (k *KafkaProcessor) confirmTransaction(transaction *appmodel.Transaction, tranactionUseCase usecase.TransactionUseCase) error {
	confirmedTransaction, err := transactionUseCase.Confirm(transaction.ID)

	if err != nil {
		return nil
	}

	topic := "bank"+confirmedTransaction.AccountFrom.Bank.Code
	transactionJson, err := transaction.ToJson()

	if err != nil {
		return nil
	}

	err = Publish(string(transactionJson), topic, k.Producer, k.deliveryChan)

	if err != nil {
		return nil
	}

	return nil
}