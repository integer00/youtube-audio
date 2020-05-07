/*

 1) Shows playlists of asked userid


get list of playlists by id (find way to lookup by username)
show results in table
get list of videos from playlist

download video(audio) into file


 USAGE:
   youtube-audio --playlists <user id>
   youtube-audio --download <url>

*/

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var (
	userpl      = flag.String("userpl", "", "user playlist")
	yotubeToken = os.Getenv("YOUTUBE_API_KEY")
)

func getPlaylists(pl string) {
	ctx := context.Background()

	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(yotubeToken))
	if err != nil {
		panic("cannot create service")
	}
	response := youtubeService.Playlists.List("snippet")
	response.ChannelId(pl)
	response.MaxResults(15)
	call, err := response.Do()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Print("stub\n")

	for _, playlist := range call.Items {
		playlistId := playlist.Id
		playlistTitle := playlist.Snippet.Title

		// Print the playlist ID and title for the playlist resource.
		fmt.Println(playlistTitle, ": ", playlistId)
	}

}

func init() {
	flag.Parse()
}

func main() {
	if yotubeToken == "" {
		log.Fatal("No YOUTUBE_API_KEY environment variable found,  extiting.")
	}

	getPlaylists(*userpl)
}
