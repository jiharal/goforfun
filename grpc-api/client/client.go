package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/s1gu/goforfun/grpc-api/customer"
)

var (
	tpl *template.Template
)

func init() {
	tpl = template.Must(template.ParseGlob("template/*.html"))
}

const (
	address = "localhost:50051"
)

func createCustomer(client pb.CustomerClient, customer *pb.CustomerRequest) {
	resp, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	if resp.Success {
		log.Printf("A new Customer has been added with id: %v", resp.Id)
	}
}

func getCustomers(client pb.CustomerClient, filter *pb.CustomerFilter) {
	// calling the streaming API
	stream, err := client.GetCustomer(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get customers: %v", err)
	}
	for {
		customer, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetCustomers(_) = _, %v", client, err)
		}
		log.Printf("Customer: %v", customer)
	}
}
func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect :%v", err)
	}
	defer conn.Close()

	http.HandleFunc("/input", func(w http.ResponseWriter, r *http.Request) {
		if err := tpl.ExecuteTemplate(w, "input.html", nil); err != nil {
			log.Fatalln("template didn't execute register: ", err)
		}
	})

	http.HandleFunc("/simpaninputan", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			email := r.FormValue("email")
			phone := r.FormValue("phone")
			client := pb.NewCustomerClient(conn)
			customer := &pb.CustomerRequest{
				Id:    123,
				Name:  name,
				Email: email,
				Phone: phone,
				Addresses: []*pb.CustomerRequest_Address{
					&pb.CustomerRequest_Address{
						Street:            "Jl. Celeban UH3/543",
						City:              "Yogyakarta",
						State:             "DIY",
						Zip:               "55198",
						IsShippingAddress: true,
					},
				},
			}
			createCustomer(client, customer)
		}
		http.RedirectHandler("localhost:9090/tampilkandata", 200)
	})

	http.HandleFunc("/tampilkandata", func(w http.ResponseWriter, r *http.Request) {
		client := pb.NewCustomerClient(conn)
		filter := &pb.CustomerFilter{Keyword: ""}
		stream, err := client.GetCustomer(context.Background(), filter)
		if err != nil {
			log.Fatalf("Error on get customers: %v", err)
		}
		for {
			customer, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("%v.GetCustomers(_) = _, %v", client, err)
			}
			if err := tpl.ExecuteTemplate(w, "tampilan.html", customer); err != nil {
				log.Fatalln("template didn't execute register: ", err)
			}
		}
	})

	http.ListenAndServe(":9090", nil)
}
