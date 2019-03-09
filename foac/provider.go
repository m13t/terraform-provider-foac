package foac

import (
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ConfigureFunc: providerConfigure,
		DataSourcesMap: map[string]*schema.Resource{
			"foac_anyway":      generateDataSource("anyway", []string{"company", "from"}),
			"foac_asshole":     generateDataSource("asshole", []string{"from"}),
			"foac_awesome":     generateDataSource("awesome", []string{"from"}),
			"foac_back":        generateDataSource("back", []string{"name", "from"}),
			"foac_bag":         generateDataSource("bag", []string{"from"}),
			"foac_ballmer":     generateDataSource("ballmer", []string{"name", "company", "from"}),
			"foac_bday":        generateDataSource("bday", []string{"name", "from"}),
			"foac_because":     generateDataSource("because", []string{"from"}),
			"foac_blackadder":  generateDataSource("blackadder", []string{"name", "from"}),
			"foac_bm":          generateDataSource("bm", []string{"name", "from"}),
			"foac_bucket":      generateDataSource("bucket", []string{"from"}),
			"foac_bus":         generateDataSource("bus", []string{"name", "from"}),
			"foac_bye":         generateDataSource("bye", []string{"from"}),
			"foac_caniuse":     generateDataSource("caniuse", []string{"tool", "from"}),
			"foac_chainsaw":    generateDataSource("chainsaw", []string{"name", "from"}),
			"foac_cocksplat":   generateDataSource("cocksplat", []string{"name", "from"}),
			"foac_cool":        generateDataSource("cool", []string{"from"}),
			"foac_cup":         generateDataSource("cup", []string{"from"}),
			"foac_dalton":      generateDataSource("dalton", []string{"name", "from"}),
			"foac_deraadt":     generateDataSource("deraadt", []string{"name", "from"}),
			"foac_diabetes":    generateDataSource("diabetes", []string{"from"}),
			"foac_donut":       generateDataSource("donut", []string{"name", "from"}),
			"foac_dosomething": generateDataSource("dosomething", []string{"do", "something", "from"}),
			"foac_equity":      generateDataSource("equity", []string{"name", "from"}),
			"foac_everyone":    generateDataSource("everyone", []string{"from"}),
			"foac_everything":  generateDataSource("everything", []string{"from"}),
			"foac_family":      generateDataSource("family", []string{"from"}),
			"foac_fascinating": generateDataSource("fascinating", []string{"from"}),
			"foac_field":       generateDataSource("field", []string{"name", "from", "reference"}),
			"foac_flying":      generateDataSource("flying", []string{"from"}),
			"foac_fts":         generateDataSource("fts", []string{"name", "from"}),
			"foac_fyyff":       generateDataSource("fyyff", []string{"from"}),
			"foac_gfy":         generateDataSource("gfy", []string{"name", "from"}),
			"foac_give":        generateDataSource("give", []string{"from"}),
			"foac_greed":       generateDataSource("greed", []string{"noun", "from"}),
			"foac_horse":       generateDataSource("horse", []string{"from"}),
			"foac_immensity":   generateDataSource("immensity", []string{"from"}),
			"foac_ing":         generateDataSource("ing", []string{"name", "from"}),
			"foac_keep":        generateDataSource("keep", []string{"name", "from"}),
			"foac_keep_calm":   generateDataSource("keepcalm", []string{"reaction", "from"}),
			"foac_king":        generateDataSource("king", []string{"name", "from"}),
			"foac_life":        generateDataSource("life", []string{"from"}),
			"foac_linus":       generateDataSource("linus", []string{"name", "from"}),
			"foac_look":        generateDataSource("look", []string{"name", "from"}),
			"foac_looking":     generateDataSource("looking", []string{"from"}),
			"foac_madison":     generateDataSource("madison", []string{"name", "from"}),
			"foac_maybe":       generateDataSource("maybe", []string{"from"}),
			"foac_me":          generateDataSource("me", []string{"from"}),
			"foac_mornin":      generateDataSource("morning", []string{"from"}),
			"foac_no":          generateDataSource("no", []string{"from"}),
			"foac_nugget":      generateDataSource("nugget", []string{"name", "from"}),
			"foac_off":         generateDataSource("off", []string{"name", "from"}),
			"foac_off_with":    generateDataSource("off-with", []string{"behavior", "from"}),
			"foac_outside":     generateDataSource("outside", []string{"name", "from"}),
			"foac_particular":  generateDataSource("particular", []string{"thing", "from"}),
			"foac_pink":        generateDataSource("pink", []string{"from"}),
			"foac_problem":     generateDataSource("problem", []string{"name", "from"}),
			"foac_programmer":  generateDataSource("programmer", []string{"from"}),
			"foac_pulp":        generateDataSource("pulp", []string{"language", "from"}),
			"foac_question":    generateDataSource("question", []string{"from"}),
			"foac_rats_arse":   generateDataSource("ratsarse", []string{"from"}),
			"foac_retard":      generateDataSource("retard", []string{"from"}),
			"foac_ridiculous":  generateDataSource("ridiculous", []string{"from"}),
			"foac_rtfm":        generateDataSource("rtfm", []string{"from"}),
			"foac_sake":        generateDataSource("sake", []string{"from"}),
			"foac_shakespeare": generateDataSource("shakespeare", []string{"name", "from"}),
			"foac_shit":        generateDataSource("shit", []string{"from"}),
			"foac_shut_up":     generateDataSource("shutup", []string{"name", "from"}),
			"foac_single":      generateDataSource("single", []string{"from"}),
			"foac_thanks":      generateDataSource("thanks", []string{"from"}),
			"foac_that":        generateDataSource("that", []string{"from"}),
			"foac_think":       generateDataSource("think", []string{"name", "from"}),
			"foac_thinking":    generateDataSource("thinking", []string{"name", "from"}),
			"foac_this":        generateDataSource("this", []string{"from"}),
			"foac_thumbs":      generateDataSource("thumbs", []string{"name", "from"}),
			"foac_too":         generateDataSource("too", []string{"from"}),
			"foac_tucker":      generateDataSource("tucker", []string{"from"}),
			"foac_what":        generateDataSource("what", []string{"from"}),
			"foac_xmas":        generateDataSource("xmas", []string{"name", "from"}),
			"foac_yoda":        generateDataSource("yoda", []string{"name", "from"}),
			"foac_you":         generateDataSource("you", []string{"name", "from"}),
			"foac_zayn":        generateDataSource("zayn", []string{"from"}),
			"foac_zero":        generateDataSource("zero", []string{"from"}),
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &client{
		http: &http.Client{},
	}

	return config, nil
}
