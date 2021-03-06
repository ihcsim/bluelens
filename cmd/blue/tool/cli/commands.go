package cli

import (
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/ihcsim/bluelens/cmd/blue/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	// CreateMusicCommand is the command line data structure for the create action of music
	CreateMusicCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// ListMusicCommand is the command line data structure for the list action of music
	ListMusicCommand struct {
		Limit       int
		Offset      int
		PrettyPrint bool
	}

	// ShowMusicCommand is the command line data structure for the show action of music
	ShowMusicCommand struct {
		ID          string
		PrettyPrint bool
	}

	// RecommendRecommendationsCommand is the command line data structure for the recommend action of recommendations
	RecommendRecommendationsCommand struct {
		// Maximum number of recommendations to be returned to the user. Set to zero to use server's default.
		Limit int
		// ID of the user these recommendations are meant for.
		UserID      string
		PrettyPrint bool
	}

	// CreateUserCommand is the command line data structure for the create action of user
	CreateUserCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// FollowUserCommand is the command line data structure for the follow action of user
	FollowUserCommand struct {
		Payload     string
		ContentType string
		FolloweeID  string
		ID          string
		PrettyPrint bool
	}

	// ListUserCommand is the command line data structure for the list action of user
	ListUserCommand struct {
		Limit       int
		Offset      int
		PrettyPrint bool
	}

	// ListenUserCommand is the command line data structure for the listen action of user
	ListenUserCommand struct {
		Payload     string
		ContentType string
		ID          string
		MusicID     string
		PrettyPrint bool
	}

	// ShowUserCommand is the command line data structure for the show action of user
	ShowUserCommand struct {
		ID          string
		PrettyPrint bool
	}

	// DownloadCommand is the command line data structure for the download command.
	DownloadCommand struct {
		// OutFile is the path to the download output file.
		OutFile string
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "create",
		Short: `create action`,
	}
	tmp1 := new(CreateMusicCommand)
	sub = &cobra.Command{
		Use:   `music ["/bluelens/music"]`,
		Short: ``,
		Long: `

Payload example:

{
   "id": "Exercitationem architecto dolorum quis dignissimos odio.",
   "tags": [
      "Cumque dolorum dolorem voluptas eveniet."
   ]
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp2 := new(CreateUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/bluelens/user"]`,
		Short: ``,
		Long: `

Payload example:

{
   "followees": [
      {
         "followees": [
            {
               "history": [
                  {
                     "id": "Exercitationem architecto dolorum quis dignissimos odio.",
                     "tags": [
                        "Cumque dolorum dolorem voluptas eveniet."
                     ]
                  }
               ],
               "id": "Iste saepe quam."
            },
            {
               "history": [
                  {
                     "id": "Exercitationem architecto dolorum quis dignissimos odio.",
                     "tags": [
                        "Cumque dolorum dolorem voluptas eveniet."
                     ]
                  }
               ],
               "id": "Iste saepe quam."
            },
            {
               "history": [
                  {
                     "id": "Exercitationem architecto dolorum quis dignissimos odio.",
                     "tags": [
                        "Cumque dolorum dolorem voluptas eveniet."
                     ]
                  }
               ],
               "id": "Iste saepe quam."
            }
         ],
         "history": [
            {
               "id": "Exercitationem architecto dolorum quis dignissimos odio.",
               "tags": [
                  "Cumque dolorum dolorem voluptas eveniet."
               ]
            }
         ],
         "id": "Iste saepe quam."
      },
      {
         "followees": [
            {
               "history": [
                  {
                     "id": "Exercitationem architecto dolorum quis dignissimos odio.",
                     "tags": [
                        "Cumque dolorum dolorem voluptas eveniet."
                     ]
                  }
               ],
               "id": "Iste saepe quam."
            },
            {
               "history": [
                  {
                     "id": "Exercitationem architecto dolorum quis dignissimos odio.",
                     "tags": [
                        "Cumque dolorum dolorem voluptas eveniet."
                     ]
                  }
               ],
               "id": "Iste saepe quam."
            },
            {
               "history": [
                  {
                     "id": "Exercitationem architecto dolorum quis dignissimos odio.",
                     "tags": [
                        "Cumque dolorum dolorem voluptas eveniet."
                     ]
                  }
               ],
               "id": "Iste saepe quam."
            }
         ],
         "history": [
            {
               "id": "Exercitationem architecto dolorum quis dignissimos odio.",
               "tags": [
                  "Cumque dolorum dolorem voluptas eveniet."
               ]
            }
         ],
         "id": "Iste saepe quam."
      }
   ],
   "history": [
      {
         "id": "Exercitationem architecto dolorum quis dignissimos odio.",
         "tags": [
            "Cumque dolorum dolorem voluptas eveniet."
         ]
      }
   ],
   "id": "Iste saepe quam."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "follow",
		Short: `Update a user's followees list with a new followee.`,
	}
	tmp3 := new(FollowUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/bluelens/user/ID/follows/FOLLOWEEID"]`,
		Short: ``,
		Long: `

Payload example:

{
   "followeeID": "Praesentium aperiam magni."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `list action`,
	}
	tmp4 := new(ListMusicCommand)
	sub = &cobra.Command{
		Use:   `music ["/bluelens/music"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp5 := new(ListUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/bluelens/user"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "listen",
		Short: `Add a music to a user's history.`,
	}
	tmp6 := new(ListenUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/bluelens/user/ID/listen/MUSICID"]`,
		Short: ``,
		Long: `

Payload example:

{
   "musicID": "Qui non rerum ullam velit."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "recommend",
		Short: `Make music recommendations for a user.`,
	}
	tmp7 := new(RecommendRecommendationsCommand)
	sub = &cobra.Command{
		Use:   `recommendations ["/bluelens/recommendations/USERID/LIMIT"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp7.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: `show action`,
	}
	tmp8 := new(ShowMusicCommand)
	sub = &cobra.Command{
		Use:   `music ["/bluelens/music/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp8.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp9 := new(ShowUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/bluelens/user/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp9.Run(c, args) },
	}
	tmp9.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp9.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)

	dl := new(DownloadCommand)
	dlc := &cobra.Command{
		Use:   "download [PATH]",
		Short: "Download file with given path",
		RunE: func(cmd *cobra.Command, args []string) error {
			return dl.Run(c, args)
		},
	}
	dlc.Flags().StringVar(&dl.OutFile, "out", "", "Output file")
	app.AddCommand(dlc)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run downloads files with given paths.
func (cmd *DownloadCommand) Run(c *client.Client, args []string) error {
	var (
		fnf func(context.Context, string) (int64, error)
		fnd func(context.Context, string, string) (int64, error)

		rpath   = args[0]
		outfile = cmd.OutFile
		logger  = goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
		ctx     = goa.WithLogger(context.Background(), logger)
		err     error
	)

	if rpath[0] != '/' {
		rpath = "/" + rpath
	}
	if rpath == "/bluelens/swagger.json" {
		fnf = c.DownloadSwaggerJSON
		if outfile == "" {
			outfile = "swagger.json"
		}
		goto found
	}
	if rpath == "/bluelens/swagger.yaml" {
		fnf = c.DownloadSwaggerYaml
		if outfile == "" {
			outfile = "swagger.yaml"
		}
		goto found
	}
	return fmt.Errorf("don't know how to download %s", rpath)
found:
	ctx = goa.WithLogContext(ctx, "file", outfile)
	if fnf != nil {
		_, err = fnf(ctx, outfile)
	} else {
		_, err = fnd(ctx, rpath, outfile)
	}
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	return nil
}

// Run makes the HTTP request corresponding to the CreateMusicCommand command.
func (cmd *CreateMusicCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/bluelens/music"
	}
	var payload client.Music
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateMusic(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateMusicCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the ListMusicCommand command.
func (cmd *ListMusicCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/bluelens/music"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListMusic(ctx, path, intFlagVal("limit", cmd.Limit), intFlagVal("offset", cmd.Offset))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListMusicCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().IntVar(&cmd.Limit, "limit", 20, ``)
	var offset int
	cc.Flags().IntVar(&cmd.Offset, "offset", offset, ``)
}

// Run makes the HTTP request corresponding to the ShowMusicCommand command.
func (cmd *ShowMusicCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/bluelens/music/%v", url.QueryEscape(cmd.ID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowMusic(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowMusicCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
}

// Run makes the HTTP request corresponding to the RecommendRecommendationsCommand command.
func (cmd *RecommendRecommendationsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/bluelens/recommendations/%v/%v", url.QueryEscape(cmd.UserID), cmd.Limit)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.RecommendRecommendations(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *RecommendRecommendationsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var limit int
	cc.Flags().IntVar(&cmd.Limit, "limit", limit, `Maximum number of recommendations to be returned to the user. Set to zero to use server's default.`)
	var userID string
	cc.Flags().StringVar(&cmd.UserID, "userID", userID, `ID of the user these recommendations are meant for.`)
}

// Run makes the HTTP request corresponding to the CreateUserCommand command.
func (cmd *CreateUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/bluelens/user"
	}
	var payload client.User
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateUser(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the FollowUserCommand command.
func (cmd *FollowUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/bluelens/user/%v/follows/%v", url.QueryEscape(cmd.ID), url.QueryEscape(cmd.FolloweeID))
	}
	var payload client.FollowUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.FollowUser(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *FollowUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var followeeID string
	cc.Flags().StringVar(&cmd.FolloweeID, "followeeID", followeeID, ``)
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
}

// Run makes the HTTP request corresponding to the ListUserCommand command.
func (cmd *ListUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/bluelens/user"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListUser(ctx, path, intFlagVal("limit", cmd.Limit), intFlagVal("offset", cmd.Offset))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().IntVar(&cmd.Limit, "limit", 20, ``)
	var offset int
	cc.Flags().IntVar(&cmd.Offset, "offset", offset, ``)
}

// Run makes the HTTP request corresponding to the ListenUserCommand command.
func (cmd *ListenUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/bluelens/user/%v/listen/%v", url.QueryEscape(cmd.ID), url.QueryEscape(cmd.MusicID))
	}
	var payload client.ListenUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListenUser(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListenUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
	var musicID string
	cc.Flags().StringVar(&cmd.MusicID, "musicID", musicID, ``)
}

// Run makes the HTTP request corresponding to the ShowUserCommand command.
func (cmd *ShowUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/bluelens/user/%v", url.QueryEscape(cmd.ID))
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowUser(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
}
