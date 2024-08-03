package usecase_test

import (
	"RestApiBackend/internal/features/posts"
	"RestApiBackend/internal/features/posts/dto"
	"RestApiBackend/internal/features/posts/repository"
	"RestApiBackend/internal/features/posts/usecase"
	"RestApiBackend/testhelpers"
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

type PostUseCaseTestSuite struct {
	suite.Suite
	pgContainer *testhelpers.PostgresContainer
	postUseCase posts.UseCase
	ctx         context.Context
	userId      uuid.UUID
}

func TestNewPostUseCase(t *testing.T) {
	suite.Run(t, new(PostUseCaseTestSuite))
}

func (suite *PostUseCaseTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	connection, err := testhelpers.CreatePostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}
	sqlConnection, _ := sql.Open("postgres", connection.ConnectionString)
	postsRepo := repository.NewPostsRepository(sqlConnection)
	postUsecase := usecase.NewPostUseCase(postsRepo)
	suite.pgContainer = connection
	suite.postUseCase = postUsecase
	suite.userId = uuid.New()
}

func (suite *PostUseCaseTestSuite) TestACreatePost0() {

	t := suite.T()
	title := "Sample Title"
	bodyText := "This is the body text of the post, which should be more than 10 characters long."
	tags := "golang,programming"

	post, err := suite.postUseCase.CreatePost(suite.ctx, suite.userId, dto.CreatePostRequest{
		Title:    &title,
		BodyText: &bodyText,
		Tags:     &tags,
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	assert.NotNil(t, post)
}

func (suite *PostUseCaseTestSuite) TestBFetchUserPosts1() {
	t := suite.T()

	p, err := suite.postUseCase.FetchPosts(suite.ctx, suite.userId)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, p.Data)
	assert.Equal(t, 1, len(p.Data))
}

func (suite *PostUseCaseTestSuite) TestCFetchAndUpdatePostAndDelete() {
	t := suite.T()
	var firstPost *dto.PostResponse
	title := "Sample Title UpdTED"

	t.Run("Fetch User's Posts", func(t *testing.T) {
		fetchPosts, err := suite.postUseCase.FetchPosts(suite.ctx, suite.userId)
		if err != nil {
			log.Fatal(err)
		}
		firstPost = fetchPosts.Data[0] // Correctly assign the pointer here
		assert.NotNil(t, firstPost)
	})

	t.Run("Update fetched post", func(t *testing.T) {
		if firstPost == nil {
			t.Fatal("firstPost is nil")
		}

		err := suite.postUseCase.UpdatePost(suite.ctx, suite.userId, firstPost.ID, dto.UpdatePostRequest{
			Title:    &title,
			BodyText: &firstPost.BodyText,
		})
		if err != nil {
			t.Fatal(err)
		}

		post, err := suite.postUseCase.FetchPost(suite.ctx, suite.userId, firstPost.ID)
		if err != nil {
			t.Fatal(err)
		}

		assert.NotNil(t, post)
		assert.Equal(t, title, post.Title)
	})

	t.Run("Delete updated post", func(t *testing.T) {
		err := suite.postUseCase.DeletePost(suite.ctx, suite.userId, firstPost.ID)
		if err != nil {
			t.Fatal(err)
		}

		assert.Nil(t, err)

		p, err := suite.postUseCase.FetchPosts(suite.ctx, suite.userId)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 0, len(p.Data))
	})
}

func (suite *PostUseCaseTestSuite) TearDownSuite() {
	log.Println("Tearing down")
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}
