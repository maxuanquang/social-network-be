package authen_and_post_svc

import (
	"context"

	"github.com/maxuanquang/social-network/internal/pkg/types"
	pb_aap "github.com/maxuanquang/social-network/pkg/types/proto/pb/authen_and_post"
)

func (a *AuthenticateAndPostService) GetUserFollower(ctx context.Context, info *pb_aap.GetUserFollowerRequest) (*pb_aap.GetUserFollowerResponse, error) {
	exist, _ := a.findUserById(info.GetUserId())
	if !exist {
		return &pb_aap.GetUserFollowerResponse{
			Status: pb_aap.GetUserFollowerResponse_USER_NOT_FOUND,
		}, nil
	}

	var user types.User
	result := a.db.Preload("Followers").First(&user, info.GetUserId())
	if result.Error != nil {
		return nil, result.Error
	}

	var followersIds []int64
	for _, follower := range user.Followers {
		followersIds = append(followersIds, int64(follower.ID))
	}
	return &pb_aap.GetUserFollowerResponse{
		Status:       pb_aap.GetUserFollowerResponse_OK,
		FollowersIds: followersIds,
	}, nil
}

func (a *AuthenticateAndPostService) GetUserFollowing(ctx context.Context, info *pb_aap.GetUserFollowingRequest) (*pb_aap.GetUserFollowingResponse, error) {
	exist, _ := a.findUserById(info.GetUserId())
	if !exist {
		return &pb_aap.GetUserFollowingResponse{
			Status: pb_aap.GetUserFollowingResponse_USER_NOT_FOUND,
		}, nil
	}

	var user types.User
	result := a.db.Preload("Followings").First(&user, info.GetUserId())
	if result.Error != nil {
		return nil, result.Error
	}

	var followingsIds []int64
	for _, following := range user.Followings {
		followingsIds = append(followingsIds, int64(following.ID))
	}
	return &pb_aap.GetUserFollowingResponse{
		Status:        pb_aap.GetUserFollowingResponse_OK,
		FollowingsIds: followingsIds,
	}, nil
}

func (a *AuthenticateAndPostService) FollowUser(ctx context.Context, info *pb_aap.FollowUserRequest) (*pb_aap.FollowUserResponse, error) {
	// Check if the user exists
	exist, _ := a.findUserById(info.GetUserId())
	if !exist {
		return &pb_aap.FollowUserResponse{Status: pb_aap.FollowUserResponse_USER_NOT_FOUND}, nil
	}
	exist, friend := a.findUserById(info.GetFollowingId())
	if !exist {
		return &pb_aap.FollowUserResponse{Status: pb_aap.FollowUserResponse_USER_NOT_FOUND}, nil
	}

	var user types.User
	a.db.Preload("Followings").First(&user, info.GetUserId())
	for _, following := range user.Followings {
		if following.ID == uint(info.GetFollowingId()) {
			return &pb_aap.FollowUserResponse{Status: pb_aap.FollowUserResponse_ALREADY_FOLLOWED}, nil
		}
	}

	err := a.db.Model(&user).Association("Followings").Append(&friend)
	if err != nil {
		return nil, err
	}
	return &pb_aap.FollowUserResponse{
		Status: pb_aap.FollowUserResponse_OK,
	}, nil
}

func (a *AuthenticateAndPostService) UnfollowUser(ctx context.Context, info *pb_aap.UnfollowUserRequest) (*pb_aap.UnfollowUserResponse, error) {
	exist, _ := a.findUserById(info.GetUserId())
	if !exist {
		return &pb_aap.UnfollowUserResponse{Status: pb_aap.UnfollowUserResponse_USER_NOT_FOUND}, nil
	}
	exist, friend := a.findUserById(info.GetFollowingId())
	if !exist {
		return &pb_aap.UnfollowUserResponse{Status: pb_aap.UnfollowUserResponse_USER_NOT_FOUND}, nil
	}

	var user types.User
	a.db.Preload("Followings").First(&user, info.GetUserId())
	currentlyFollowing := false
	for _, following := range user.Followings {
		if following.ID == uint(info.GetFollowingId()) {
			currentlyFollowing = true
			break
		}
	}
	if !currentlyFollowing {
		return &pb_aap.UnfollowUserResponse{Status: pb_aap.UnfollowUserResponse_NOT_FOLLOWED}, nil
	}

	err := a.db.Model(&user).Association("Followings").Delete(&friend)
	if err != nil {
		return nil, err
	}
	return &pb_aap.UnfollowUserResponse{
		Status: pb_aap.UnfollowUserResponse_OK,
	}, nil
}

func (a *AuthenticateAndPostService) GetUserPosts(ctx context.Context, info *pb_aap.GetUserPostsRequest) (*pb_aap.GetUserPostsResponse, error) {
	exist, _ := a.findUserById(info.GetUserId())
	if !exist {
		return &pb_aap.GetUserPostsResponse{Status: pb_aap.GetUserPostsResponse_USER_NOT_FOUND}, nil
	}

	var user types.User
	a.db.Preload("Posts").First(&user, info.GetUserId())

	// Return
	var posts_ids []int64
	for _, post := range user.Posts {
		posts_ids = append(posts_ids, int64(post.ID))
	}

	return &pb_aap.GetUserPostsResponse{
		Status:   pb_aap.GetUserPostsResponse_OK,
		PostsIds: posts_ids,
	}, nil
}
