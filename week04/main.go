package main

func main() {
	Constructor()
}

//寻找旋转排序数组中的最小值 II
func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			//如果下标为mid对应的值比最后一位的值要小，说明最小值在mid的左边，或者就是当前下标为mid的值，将mid定为右边界
			right = mid
		} else if nums[mid] > nums[right] {
			//如果下标为mid对应的值比最后一位的值要大，说明最小值一定在mid的右边，将mid右移一位的定位左边界
			left = mid + 1
		} else {
			//当存在相等的情况，无法判断，就将右边界往左移一位
			right--
		}
	}
	return nums[right]
}

//设计推特
type Twitter struct {
	Tweets  [][]int       //存储所有推特内容，例：[[1,5],[2,6]];[1,5]中，1表示userId，5表示推特记录ID
	Follows map[int][]int //存储用户关注列表，key为userId，value为关注的userId数组
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	twitter := Twitter{
		Tweets:  [][]int{},
		Follows: make(map[int][]int),
	}
	return twitter
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	var content []int
	content = append(content, userId)
	content = append(content, tweetId)
	this.Tweets = append(this.Tweets, content)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	//获取当前用户的关注列表以及自己
	users := this.Follows[userId]
	users = append(users, userId)
	var result []int
	if len(users) > 0 {
		for j := len(this.Tweets) - 1; j >= 0; j-- {
			tweet := this.Tweets[j]
			for i := 0; i < len(users); i++ {
				if tweet[0] == users[i] {
					//只找出最近十条记录
					if len(result) == 10 {
						return result
					}
					result = append(result, tweet[1])
					break
				}
			}
		}
	}
	return result
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	followers, exist := this.Follows[followerId]
	if exist {
		for i := 0; i < len(followers); i++ {
			//已经关注的用户，不再重复关注
			if followers[i] == followerId {
				return
			}
		}
	}
	followers = append(followers, followeeId)
	this.Follows[followerId] = followers
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	followers, exist := this.Follows[followerId]
	if exist {
		var newFollowers []int
		for i := 0; i < len(followers); i++ {
			if followers[i] == followeeId {
				continue
			} else {
				newFollowers = append(newFollowers, followers[i])
			}
		}
		this.Follows[followerId] = newFollowers
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
