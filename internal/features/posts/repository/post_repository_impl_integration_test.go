package repository_test

import (
	"RestApiBackend/internal/features/posts"
	"RestApiBackend/internal/features/posts/entites"
	"RestApiBackend/internal/features/posts/repository"
	"RestApiBackend/testhelpers"
	"context"
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

type PostRepositoryIntegrationTestSuite struct {
	suite.Suite
	pgContainer *testhelpers.PostgresContainer
	repository  posts.PostRepository
	ctx         context.Context
	userId      uuid.UUID
}

func TestNewPostsRepository(t *testing.T) {
	suite.Run(t, new(PostRepositoryIntegrationTestSuite))
}

func (suite *PostRepositoryIntegrationTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	connection, err := testhelpers.CreatePostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}
	sqlConnection, err := sql.Open("postgres", connection.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	postsRepo := repository.NewPostsRepository(sqlConnection)
	suite.pgContainer = connection
	suite.repository = postsRepo
	suite.userId = uuid.New()
}

func (suite *PostRepositoryIntegrationTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *PostRepositoryIntegrationTestSuite) TestFetchPost() {
	t := suite.T()
	var createdPost *entites.Post
	t.Run("Create Post", func(t *testing.T) {
		creatingPost := entites.NewPost()
		creatingPost.Title = "Test post"
		creatingPost.Body = "Body test"
		created, err := suite.repository.CreatePostForUser(suite.ctx, suite.userId, *creatingPost)
		if err != nil {
			t.Fatal(err)
		}
		createdPost = created
		assert.NotNil(t, createdPost)
		assert.NotNil(t, createdPost.ID)
	})

	t.Run("Fetch created post by id", func(t *testing.T) {
		p, err := suite.repository.FetchPost(suite.ctx, suite.userId, createdPost.ID)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotNil(t, p)
		assert.Equal(t, "Test post", p.Title)
	})
}
