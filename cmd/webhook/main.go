package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eddielau42/lalamove-go-api/enum"
	"github.com/eddielau42/lalamove-go-api/lalamove"
)

const (
	envPrefix = "LALAMOVE_"
)

var (
	cli *lalamove.Client

	// apikey, secret, market
	apikey, secret, market string

	// webhook地址
	webhookURL string

	isSandbox bool
)

func init() {
	loadEnv("")
}

func main() {
	flag.StringVar(&apikey, "apikey", "", "apikey")
	flag.StringVar(&secret, "secret", "", "secret")
	flag.StringVar(&market, "market", "", "地区")
	flag.StringVar(&webhookURL, "url", "", "要设置webhook的地址")
	
	flag.Parse()
	
	if apikey == "" {
		fmt.Println("请输入apikey!")
		return
	}
	if secret == "" {
		fmt.Println("请输入secret!")
		return
	}
	if webhookURL == "" {
		fmt.Println("请输入要设置webhook地址!")
		return
	}
	if market == "" {
		market = enum.COUNTRY_HONGKONG
	}

	// Check sandbox
	isSandbox = !strings.Contains(apikey, "pk_prod") || !strings.Contains(secret, "sk_prod")

	handle()
}

func handle() {
	fmt.Printf(">>> 开始设置webhook地址...\n")

	// conf := getLalamoveConfigFromEnv()
	conf := lalamove.Config{
		Apikey: apikey,
		Secret: secret,
		Country: market,
	}

	cli = lalamove.NewClient(conf)
	if isSandbox {
		cli.Sandbox()
	}

	ok, err := cli.SetWebhook(webhookURL)
	if err != nil {
		fmt.Println("----- " + err.Error())
	}

	if ok {
		fmt.Println("<<< 设置成功。")
		return
	}

	fmt.Println(`<<< 设置失败！(详情请参考官方文档: https://developers.lalamove.com/?shell#webhook)`)
}

// getLalamoveConfigFromEnv 通过读取环境配置, 返回lalamove实例配置信息
func getLalamoveConfigFromEnv() lalamove.Config {
	country := os.Getenv("LALAMOVE_MARKET")
	if country == "" {
		country = enum.COUNTRY_HONGKONG
	}

	conf := lalamove.Config{
		Country: country,
	}

	isSandbox, _ = strconv.ParseBool(os.Getenv("LALAMOVE_SANDBOX"))
	if isSandbox {
		conf.Apikey = os.Getenv("LALAMOVE_SANDBOX_APIKEY")
		conf.Secret = os.Getenv("LALAMOVE_SANDBOX_SECRET")
	} else {
		conf.Apikey = os.Getenv("LALAMOVE_APIKEY")
		conf.Secret = os.Getenv("LALAMOVE_SECRET")
	}

	return conf
}

// loadEnv 读取本地环境变量配置文件
func loadEnv(envDir string) error {
	if envDir == "" {
		envDir = ".env"
	}

	file, err := os.Open(envDir)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		row, _, err := reader.ReadLine()
		if len(row) > 0 {
			s := strings.Split(string(row), "=")
			if len(s[0]) > 0 && len(s[1]) > 0 {
				key := strings.ToUpper(envPrefix + strings.Trim(strings.TrimPrefix(s[0], envPrefix), " "))
                val := strings.Trim(strings.Trim(strings.Trim(s[1], "'"), "\""), " ")
                // 写入系统环境变量
                os.Setenv(key, val)
			}
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
			break
		}
	}
	
	return nil
}