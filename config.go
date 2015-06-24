// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package flydb

import (
    "path/filepath"
)

type Config struct {
    Path string
    FormatName string
    Format Format
    Save bool
    // save database every `SaveInterval` milliseconds
    SaveInterval int
}

func setDefaultConfig(config Config) Config {
    // format
    var format Format
    if config.FormatName != "" {
        format = GuessFormat(config.FormatName)
    } else if config.Format != nil {
        format = config.Format
    } else {
        ext := filepath.Ext(config.Path)
        format = CheckFormatByExtension(ext)
    }

    return Config {
        Path: config.Path,
        FormatName: config.FormatName,
        Format: format,
        Save: config.Save,
        SaveInterval: config.SaveInterval,
    }
}