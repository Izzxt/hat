/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sync"
	"time"

	"github.com/Izzxt/hat/client"
	"github.com/Izzxt/hat/downloader"
	"github.com/Izzxt/hat/fs"
	"github.com/spf13/cobra"
)

var badgeName string

// badgepartsCmd represents the bagdeparts command
var badgepartsCmd = &cobra.Command{
	Use:   "badgeparts",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup

		c := client.NewClient()
		d := downloader.NewDownloader(c)
		d.SetDomain(Domain)
		d.SetOutput(Output)
		d.SetOther()
		d.SetPath("/c_images/Badgeparts/")
		if badgeName != "" {
			d.SetFileName(fmt.Sprintf("%s.png", badgeName))
			d.Download()
		} else {
			if Output != "" {
				d.SetOutput(Output)
			} else {
				d.SetOutput("resource/c_images/Badgeparts/")
			}

			badges := []string{
				"badgepart_base_basic_1.png", "badgepart_base_basic_2.png", "badgepart_base_basic_3.png", "badgepart_base_basic_4.png", "badgepart_base_basic_5.png",
				"badgepart_base_advanced_1.png", "badgepart_base_advanced_2.png", "badgepart_base_advanced_3.png", "badgepart_base_advanced_4.png", "badgepart_base_gold_1_part2.png",
				"badgepart_base_gold_1_part1.png", "badgepart_base_gold_2_part2.png", "badgepart_base_gold_2_part1.png", "badgepart_base_pin_part2.png", "badgepart_base_pin_part1.png",
				"badgepart_base_gradient_1.png", "badgepart_base_gradient_2.png", "badgepart_base_circles_1.png", "badgepart_base_circles_2.png", "badgepart_base_ornament_1_part2.png",
				"badgepart_base_ornament_1_part1.png", "badgepart_base_ornament_2_part2.png", "badgepart_base_ornament_2_part1.png", "badgepart_base_misc_1_part2.png", "badgepart_base_misc_1_part1.png",
				"badgepart_base_misc_2.png", "badgepart_base_beams_part2.png", "badgepart_base_beams_part1.png", "badgepart_base_ring.png", "badgepart_base_simplestar_part2.png",
				"badgepart_base_simplestar_part1.png", "badgepart_base_spiral.png", "badgepart_base_book.png", "badgepart_base_egg.png", "badgepart_base_ornament.png",
				"badgepart_base_shield_part2.png", "badgepart_base_shield_part1.png", "badgepart_symbol_background_1.png", "badgepart_symbol_background_2.png", "badgepart_symbol_background_3_part2.png",
				"badgepart_symbol_background_3_part1.png", "badgepart_symbol_ball_1_part2.png", "badgepart_symbol_ball_1_part1.png", "badgepart_symbol_ball_2_part2.png", "badgepart_symbol_ball_2_part1.png",
				"badgepart_symbol_bobba.png", "badgepart_symbol_bomb_part2.png", "badgepart_symbol_bomb_part1.png", "badgepart_symbol_bow.png", "badgepart_symbol_box_1.png",
				"badgepart_symbol_box_2.png", "badgepart_symbol_bunting_1.png", "badgepart_symbol_bunting_2.png", "badgepart_symbol_butterfly_part2.png", "badgepart_symbol_butterfly_part1.png",
				"badgepart_symbol_cowskull_part2.png", "badgepart_symbol_cowskull_part1.png", "badgepart_symbol_cross.png", "badgepart_symbol_diamond.png", "badgepart_symbol_diploma_part2.png",
				"badgepart_symbol_diploma_part1.png", "badgepart_symbol_eyeball_part2.png", "badgepart_symbol_eyeball_part1.png", "badgepart_symbol_fist.png", "badgepart_symbol_flame_1.png",
				"badgepart_symbol_flame_2.png", "badgepart_symbol_flash.png", "badgepart_symbol_flower_1_part2.png", "badgepart_symbol_flower_1_part1.png", "badgepart_symbol_flower_2.png",
				"badgepart_symbol_flower_3.png", "badgepart_symbol_flower_4.png", "badgepart_symbol_football.png", "badgepart_symbol_heart_1_part2.png", "badgepart_symbol_heart_1_part1.png",
				"badgepart_symbol_heart_2_part2.png", "badgepart_symbol_heart_2_part1.png", "badgepart_symbol_jingjang_part2.png", "badgepart_symbol_jingjang_part1.png", "badgepart_symbol_lips_part2.png",
				"badgepart_symbol_lips_part1.png", "badgepart_symbol_note.png", "badgepart_symbol_peace.png", "badgepart_symbol_planet_part2.png", "badgepart_symbol_planet_part1.png",
				"badgepart_symbol_rainbow_part2.png", "badgepart_symbol_rainbow_part1.png", "badgepart_symbol_rosete.png", "badgepart_symbol_shape.png", "badgepart_symbol_star_1.png",
				"badgepart_symbol_star_2.png", "badgepart_symbol_sword_1_part2.png", "badgepart_symbol_sword_1_part1.png", "badgepart_symbol_sword_2_part2.png", "badgepart_symbol_sword_2_part1.png",
				"badgepart_symbol_sword_3_part2.png", "badgepart_symbol_sword_3_part1.png", "badgepart_symbol_wings_1.png", "badgepart_symbol_wings_2.png", "badgepart_symbol_arrow_down.png",
				"badgepart_symbol_arrow_left.png", "badgepart_symbol_arrow_right.png", "badgepart_symbol_arrow_up.png", "badgepart_symbol_arrowbig_up.png", "badgepart_symbol_axe_part2.png",
				"badgepart_symbol_axe_part1.png", "badgepart_symbol_bug_part2.png", "badgepart_symbol_bug_part1.png", "badgepart_symbol_capsbig_part2.png", "badgepart_symbol_capsbig_part1.png",
				"badgepart_symbol_capssmall_part2.png", "badgepart_symbol_capssmall_part1.png", "badgepart_symbol_cloud.png", "badgepart_symbol_crown_part2.png", "badgepart_symbol_crown_part1.png",
				"badgepart_symbol_diamsmall2.png", "badgepart_symbol_diamsmall.png", "badgepart_symbol_drop.png", "badgepart_symbol_fingersheavy.png", "badgepart_symbol_fingersv.png",
				"badgepart_symbol_gtr_part2.png", "badgepart_symbol_gtr_part1.png", "badgepart_symbol_hat.png", "badgepart_symbol_oval_part2.png", "badgepart_symbol_oval_part1.png",
				"badgepart_symbol_pawprint.png", "badgepart_symbol_screw.png", "badgepart_symbol_stickL_part2.png", "badgepart_symbol_stickL_part1.png", "badgepart_symbol_stickR_part2.png",
				"badgepart_symbol_stickR_part1.png", "badgepart_symbol_alligator.png", "badgepart_symbol_americanfootball_part2.png", "badgepart_symbol_americanfootball_part1.png", "badgepart_symbol_award_part2.png",
				"badgepart_symbol_award_part1.png", "badgepart_symbol_bananapeel.png", "badgepart_symbol_battleball.png", "badgepart_symbol_biohazard.png", "badgepart_symbol_bird.png",
				"badgepart_symbol_bishop.png", "badgepart_symbol_coalion.png", "badgepart_symbol_cocoamug.png", "badgepart_symbol_dashflag.png", "badgepart_symbol_diamondring_part2.png",
				"badgepart_symbol_diamondring_part1.png", "badgepart_symbol_discoball_part2.png", "badgepart_symbol_discoball_part1.png", "badgepart_symbol_dog.png", "badgepart_symbol_electricguitarh_part2.png",
				"badgepart_symbol_electricguitarh_part1.png", "badgepart_symbol_electricguitarv_part2.png", "badgepart_symbol_electricguitarv_part1.png", "badgepart_symbol_film.png", "badgepart_symbol_flame_part2.png",
				"badgepart_symbol_flame_part1.png", "badgepart_symbol_gamepad.png", "badgepart_symbol_gem1_part2.png", "badgepart_symbol_gem1_part1.png", "badgepart_symbol_gem2_part2.png",
				"badgepart_symbol_gem2_part1.png", "badgepart_symbol_gem3_part2.png", "badgepart_symbol_gem3_part1.png", "badgepart_symbol_hawk.png", "badgepart_symbol_hearts_down.png",
				"badgepart_symbol_hearts_up.png", "badgepart_symbol_horseshoe.png", "badgepart_symbol_inksplatter.png", "badgepart_symbol_leaf.png", "badgepart_symbol_micstand.png",
				"badgepart_symbol_mirror_part2.png", "badgepart_symbol_mirror_part1.png", "badgepart_symbol_monkeywrench.png", "badgepart_symbol_note1.png", "badgepart_symbol_note2.png",
				"badgepart_symbol_note3.png", "badgepart_symbol_nursecross.png", "badgepart_symbol_pencil_part2.png", "badgepart_symbol_pencil_part1.png", "badgepart_symbol_queen.png",
				"badgepart_symbol_rock.png", "badgepart_symbol_rook.png", "badgepart_symbol_skate.png", "badgepart_symbol_smallring_part2.png", "badgepart_symbol_smallring_part1.png",
				"badgepart_symbol_snowstorm_part2.png", "badgepart_symbol_snowstorm_part1.png", "badgepart_symbol_sphere.png", "badgepart_symbol_spraycan_part2.png", "badgepart_symbol_spraycan_part1.png",
				"badgepart_symbol_stars1.png", "badgepart_symbol_stars2.png", "badgepart_symbol_stars3.png", "badgepart_symbol_stars4.png", "badgepart_symbol_stars5.png",
				"badgepart_symbol_waterdrop_part2.png", "badgepart_symbol_waterdrop_part1.png", "badgepart_symbol_wolverine.png", "badgepart_symbol_0.png", "badgepart_symbol_1.png",
				"badgepart_symbol_2.png", "badgepart_symbol_3.png", "badgepart_symbol_4.png", "badgepart_symbol_5.png", "badgepart_symbol_6.png",
				"badgepart_symbol_7.png", "badgepart_symbol_8.png", "badgepart_symbol_9.png", "badgepart_symbol_a.png", "badgepart_symbol_b.png",
				"badgepart_symbol_c.png", "badgepart_symbol_d.png", "badgepart_symbol_e.png", "badgepart_symbol_f.png", "badgepart_symbol_g.png",
				"badgepart_symbol_h.png", "badgepart_symbol_i.png", "badgepart_symbol_j.png", "badgepart_symbol_k.png", "badgepart_symbol_l.png",
				"badgepart_symbol_m.png", "badgepart_symbol_n.png", "badgepart_symbol_o.png", "badgepart_symbol_p.png", "badgepart_symbol_q.png",
				"badgepart_symbol_r.png", "badgepart_symbol_s.png", "badgepart_symbol_t.png", "badgepart_symbol_u.png", "badgepart_symbol_v.png",
				"badgepart_symbol_w.png", "badgepart_symbol_x.png", "badgepart_symbol_y.png", "badgepart_symbol_z.png", "badgepart_symbol_pixel_part2.png",
				"badgepart_symbol_pixel_part1.png", "badgepart_symbol_credit_part2.png", "badgepart_symbol_credit_part1.png", "badgepart_symbol_hc_part2.png", "badgepart_symbol_hc_part1.png",
				"badgepart_symbol_vip_part2.png", "badgepart_symbol_vip_part1.png",
			}

			for _, v := range badges {
				wg.Add(1)
				exts := fs.IsFileExists(d.GetOutput(), v)
				if !exts {
					go func(v string) {
						defer wg.Done()
						d.SetFileName(v)
						d.Download()
					}(v)
					time.Sleep(100 * time.Millisecond)
				}
			}
			wg.Wait()
		}
	},
}

func init() {
	rootCmd.AddCommand(badgepartsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	badgepartsCmd.PersistentFlags().StringVarP(&badgeName, "name", "n", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bagdepartsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
