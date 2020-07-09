package graph

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"meetmeup/graph/model"
	"net/http"
	"time"
)

const userloaderKey = "userloader"

func DataLoaderMiddleware(db *mongo.Collection, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userLoader := UserLoader{
			wait:     1 * time.Millisecond,
			maxBatch: 100,
			fetch: func(ids []string) ([]*model.User, []error) {
				var users []*model.User
				fmt.Println("ids", ids)

				// TODO: not sure of this line as i am using mongo instead of postgres
				err := db.FindOne(context.TODO(), bson.D{{"_id", ids}}).Decode(&users)

				if err != nil {
					return nil, []error{err}
				}

				u := make(map[string]*model.User, len(users))

				for _, user := range users {
					u[user.ID] = user
				}

				result := make([]*model.User, len(ids))

				for i, id := range ids {
					result[i] = u[id]
				}

				return result, []error{err}
			},
		}

		ctx := context.WithValue(r.Context(), userloaderKey, &userLoader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userloaderKey).(*UserLoader)
}
